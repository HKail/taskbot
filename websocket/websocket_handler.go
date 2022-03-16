package websocket

import "github.com/hkail/taskbot/dto"

// DefaultHandlers 默认的 websocket event handler
var DefaultHandlers struct {
	ATMessage ATMessageEventHandler
}

// RegisterHandlers 注册事件处理 handler
func RegisterHandlers(handlers ...interface{}) dto.Intent {
	intent := dto.Intent(0)

	for _, handler := range handlers {
		switch handler.(type) {
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handler.(ATMessageEventHandler)
			intent |= dto.IntentAtMessages
		// TODO 支持其它 event 注册
		default:
		}
	}

	return intent
}

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *dto.WSPayload, data *dto.WSATMessageData) error
