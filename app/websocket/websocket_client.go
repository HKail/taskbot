package websocket

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	dto2 "github.com/hkail/taskbot/app/dto"

	"github.com/gorilla/websocket"
)

const (
	defaultMessageChanSize = 10000
)

type messageChan chan *dto2.WSPayload

type closeErrorChan chan error

type WSClient struct {
	version         int
	conn            *websocket.Conn
	heartbeatTicket *time.Ticker
	session         *dto2.WSSession
	messageChan     messageChan
	closeChan       closeErrorChan
}

func NewWSClient(session *dto2.WSSession) *WSClient {
	return &WSClient{
		session:         session,
		heartbeatTicket: time.NewTicker(45 * time.Second), // 在收到 hello 包之后, 会使用其返回的心跳时间进行重置
		messageChan:     make(messageChan, defaultMessageChanSize),
		closeChan:       make(closeErrorChan, 10),
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
func (c *WSClient) SendMessage(message *dto2.WSPayload) error {
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
		log.Printf("websocket dial err=%v", err)
		return err
	}

	return nil
}

// Identify 对链接进行鉴权
func (c *WSClient) Identify() error {
	if c.session.Intent == 0 {
		return errors.New("zero is an invalid intent value")
	}

	event := &dto2.WSPayload{
		Data: &dto2.WSIdentityData{
			Token:   c.session.Token.GetString(),
			Intents: c.session.Intent,
		},
	}
	event.OPCode = dto2.OPCodeIdentify

	return c.SendMessage(event)
}

// Resume 重连
func (c *WSClient) Resume() error {
	event := &dto2.WSPayload{
		Data: &dto2.WSResumeData{
			Token: c.session.Token.GetString(),
		},
	}
	event.OPCode = dto2.OPCodeResume

	return c.SendMessage(event)
}

// Listening 已阻塞的形式开始监听 websocket 的所有事件
func (c *WSClient) Listening() error {
	// 从 websocket 中读取消息并发送至消息缓冲 chan 中
	go c.readMessageToChan()
	// 从消息缓冲 chan 中消费消息并处理
	go c.listenAndHandleMessage()

	for {
		select {
		case err := <-c.closeChan: // 连接关闭
			if websocket.IsCloseError(err, errCodeSendMessageTooFast, errCodeSessionTimeout) { // 可以直接重连
				return NewWSError(errCodeConnNeedReconnect, err.Error())
			}

			// TODO 需要处理 4900～4913 这些内部错误码
			if websocket.IsCloseError(err) { // 可以重新鉴权
				return NewWSError(errCodeConnNeedReIdentify, err.Error())
			}

			// 无法处理的错误
			return NewWSError(errCodeConnNeedPanic, err.Error())

		case <-c.heartbeatTicket.C: // 心跳维持
			heartbeatData := &dto2.WSPayload{
				WSPayloadBase: dto2.WSPayloadBase{
					OPCode: dto2.OPCodeHeartbeat,
				},
				Data: c.session.LastSeq,
			}

			err := c.SendMessage(heartbeatData)
			if err != nil {
				log.Printf("heartbeat send err=%v", err)
			}
		}
	}
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

		event := &dto2.WSPayload{}
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
		if event.Type == dto2.EventReady {
			c.readyEventHandler(event)
			continue
		}

		fmt.Println("in this")
		err := parseAndHandleEvent(event)
		if err != nil {
			log.Printf("parseAndHandleEvent has err=%v", err)
		}
	}
}

func (c *WSClient) saveSeq(seq uint32) {
	if seq > 0 {
		c.session.LastSeq = seq
	}
}

func (c *WSClient) isBuildInEventAndHandle(event *dto2.WSPayload) bool {
	switch event.OPCode {
	case dto2.OPCodeHello: // 完成连接, 需要开始维持心跳
		c.startHeartbeatTicker(event.RawMessage)
	case dto2.OPCodeReconnect: // 达到连接时长, 需要进行重连
		c.closeChan <- ErrNeedReconnect
	case dto2.OPCodeInvalidSession: // session 无效, 需要重新鉴权
		c.closeChan <- ErrInvalidSession
	case dto2.OPCodeHeartbeatACK: // 心跳 ack, 无需处理
	default:
		return false
	}

	return true
}

func (c *WSClient) startHeartbeatTicker(message []byte) {
	helloData := &dto2.WSHelloData{}
	if err := parseData(message, helloData); err != nil {
		log.Println(err)
		// TODO 是否应该提前结束呢
	}

	c.heartbeatTicket.Reset(time.Duration(helloData.HeartbeatInterval) * time.Millisecond)
}

func (c *WSClient) readyEventHandler(event *dto2.WSPayload) {
	readyData := &dto2.WSReadyData{}
	if err := parseData(event.RawMessage, readyData); err != nil {
		log.Println(err)
		// TODO 是否应该提前结束呢
	}

	c.version = readyData.Version
	c.session.ID = readyData.SessionID
	// TODO 支持分片
}
