package websocket

import (
	"fmt"
	"log"
	"runtime"

	"github.com/hkail/taskbot/dto"
	"github.com/hkail/taskbot/token"
)

// WSManager 用于启动和管理 websocket 链接
type WSManager struct {
	sessionChan chan dto.WSSession
}

func (m *WSManager) Start(wsAp *dto.WSAccessPoint, token *token.Token, intent dto.Intent) error {
	// TODO 支持分片
	m.sessionChan = make(chan dto.WSSession, 1)
	m.sessionChan <- dto.WSSession{
		URL:    wsAp.URL,
		Token:  *token,
		Intent: intent,
	}

	for session := range m.sessionChan {
		fmt.Println(session)
	}

	return nil
}

func (m *WSManager) newWSConnect(session dto.WSSession) {
	defer func() {
		// 防止进程退出以及打印错误堆栈信息
		if err := recover(); err != nil {
			panicHandler(err, session)
			// 放回 session chan 中, 以尝试重新建立 websocket 链接
			m.sessionChan <- session
		}
	}()

	wsClient := NewWSClient(&session)
	if err := wsClient.Connect(); err != nil {
		log.Println(err)
		m.sessionChan <- session // 连接失败, 丢回 session chan 进行重连
		return
	}

	var err error
	if session.ID == "" { // 初次连接
		err = wsClient.Identify()
	} else {
		err = wsClient.Resume()
	}
	if err != nil {
		log.Println(err)
		return
	}

	if err := wsClient.Listening(); err != nil {

	}
}

func panicHandler(e interface{}, session dto.WSSession) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]

	log.Printf("[PANIC]session=[%#v], err=[%v], stack=%s", session, e, buf)
}
