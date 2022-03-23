package websocket

import (
	"log"
	"runtime"

	"github.com/hkail/taskbot/app/dto"
	"github.com/hkail/taskbot/app/token"
)

// WSManager 用于启动和管理 websocket 链接
type WSManager struct {
	sessionChan chan dto.WSSession
}

func NewWSManager() *WSManager {
	return &WSManager{}
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
		go m.newWSConnect(session)
	}

	return nil
}

func (m *WSManager) newWSConnect(session dto.WSSession) {
	defer func() {
		// 防止进程退出以及打印错误堆栈信息
		if err := recover(); err != nil {
			panicHandler(err, session)
			// 放回 session chan 中, 以尝试重新建立 websocket 链接
			//m.sessionChan <- session
		}
	}()

	wsClient := NewWSClient(&session)
	if err := wsClient.Connect(); err != nil {
		log.Println(err)
		//m.sessionChan <- session // 连接失败, 丢回 session chan 进行重连
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
		curSession := wsClient.session

		if IsNeedReIdentifyError(err) { // 重新鉴权, 需要清空 session 和 lastSeq 信息
			curSession.LastSeq = 0
			curSession.ID = ""
		}

		if IsNeedPanicError(err) { // 无法重新鉴权的错误, 已经无法恢复了, 只能 panic
			log.Printf("the connect can't re-identify err=%v", err)
			panic(err)
		}

		// 将 session 发送回 session chan 中重新使用
		m.sessionChan <- *curSession
	}
}

func panicHandler(e interface{}, session dto.WSSession) {
	buf := make([]byte, 1024)
	buf = buf[:runtime.Stack(buf, false)]

	log.Printf("[PANIC]session=[%#v], err=[%v], stack=%s", session, e, buf)
}
