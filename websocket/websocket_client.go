package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hkail/taskbot/dto"

	"github.com/gorilla/websocket"
)

type messageChan chan *dto.WSPayload

type closeErrorChan chan error

type WSClient struct {
	version         int
	conn            *websocket.Conn
	heartbeatTicket *time.Ticker
	session         *dto.WSSession
	messageChan     messageChan
	closeChan       closeErrorChan
}

func NewWSClient(session *dto.WSSession) *WSClient {
	return &WSClient{
		session: session,
	}
}

// Close 关闭链接
func (c *WSClient) Close() {
	if err := c.conn.Close(); err != nil {
		log.Println(err)
	}

	c.heartbeatTicket.Stop()
}

// SendMessage 消息发送
func (c *WSClient) SendMessage(message *dto.WSPayload) error {
	// 此处必定不会出错, 因此忽略 error
	m, _ := json.Marshal(message)

	if err := c.conn.WriteMessage(websocket.TextMessage, m); err != nil {
		return err
	}

	return nil
}

// Connect 建立 websocket 链接
func (c *WSClient) Connect() error {
	var err error

	c.conn, _, err = websocket.DefaultDialer.Dial(c.session.URL, nil)
	if err != nil {
		return err
	}

	return nil
}

// Identify 对链接进行鉴权
func (c *WSClient) Identify() error {
	if c.session.Intent == 0 {
		return errors.New("zero is an invalid intent value")
	}

	event := &dto.WSPayload{
		Data: &dto.WSIdentityData{
			Token:   c.session.Token.GetString(),
			Intents: c.session.Intent,
		},
	}
	event.OPCode = dto.OPCodeIdentify

	return c.SendMessage(event)
}

// Resume 重连
func (c *WSClient) Resume() error {
	event := &dto.WSPayload{
		Data: &dto.WSResumeData{
			Token: c.session.Token.GetString(),
		},
	}
	event.OPCode = dto.OPCodeResume

	return c.SendMessage(event)
}

// Listening 已阻塞的形式开始监听 websocket 的所有事件
func (c *WSClient) Listening() error {

	return nil
}

func (c *WSClient) readMessageToChan() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			close(c.messageChan)
			c.closeChan <- err
			return
		}

		event := &dto.WSPayload{}
		if err := json.Unmarshal(message, &event); err != nil {
			log.Println(err)
			continue
		}

		event.RawMessage = message
		if c.isBuildInEventAndHandle(event) { // 判断是否为内部事件并进行处理
			continue
		}

		c.messageChan <- event
	}
}

func (c *WSClient) listenAndHandleMessage() {
	defer func() {
		if err := recover(); err != nil {
			panicHandler(err, *c.session)
			c.closeChan <- fmt.Errorf("panic: %v", err)
		}
	}()

	for event := range c.messageChan {
		c.saveSeq(event.Seq)
		// 对 ready 事件进行特殊处理
		if event.Type == "READY" {
			c.readyEventHandler(event)
			continue
		}

		// TODO 具体 event 处理逻辑
	}
}

func (c *WSClient) saveSeq(seq uint32) {
	if seq > 0 {
		c.session.LastSeq = seq
	}
}

func (c *WSClient) isBuildInEventAndHandle(event *dto.WSPayload) bool {
	switch event.OPCode {
	case dto.OPCodeHello: // 完成连接, 需要开始维持心跳
		c.startHeartbeatTicker(event.RawMessage)
	case dto.OPCodeReconnect: // 达到连接时长, 需要进行重连
		c.closeChan <- ErrNeedReconnect
	case dto.OPCodeInvalidSession: // session 无效, 需要重新鉴权
		c.closeChan <- ErrInvalidSession
	case dto.OPCodeHeartbeatACK: // 心跳 ack, 无需处理
	default:
		return false
	}

	return true
}

func (c *WSClient) startHeartbeatTicker(message []byte) {
	helloData := &dto.WSHelloData{}
	if err := parseData(message, helloData); err != nil {
		log.Println(err)
		// TODO 是否应该提前结束呢
	}

	c.heartbeatTicket.Reset(time.Duration(helloData.HeartbeatInterval) * time.Millisecond)
}

func (c *WSClient) readyEventHandler(event *dto.WSPayload) {
	readyData := &dto.WSReadyData{}
	if err := parseData(event.RawMessage, readyData); err != nil {
		log.Println(err)
		// TODO 是否应该提前结束呢
	}

	c.version = readyData.Version
	c.session.ID = readyData.SessionID
	// TODO 支持分片
}
