package websocket

import (
	dto2 "github.com/hkail/taskbot/app/dto"
)

// DefaultHandlers 默认的 websocket event handler
var DefaultHandlers struct {
	ATMessage ATMessageEventHandler
}

// RegisterHandlers 注册事件处理 handler
func RegisterHandlers(handlers ...interface{}) dto2.Intent {
	intent := dto2.Intent(0)

	for _, handler := range handlers {
		switch handler.(type) {
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handler.(ATMessageEventHandler)
			intent |= dto2.IntentAtMessages
		// TODO 支持其它 event 注册
		default:
		}
	}

	return intent
}

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *dto2.WSPayload, data *dto2.WSATMessageData) error
