package websocket

import (
	"fmt"

	dto2 "github.com/hkail/taskbot/app/dto"
)

func parseAndHandleEvent(event *dto2.WSPayload) error {
	fmt.Println(event)

	if event.OPCode == dto2.OPCodeDispatch {
		switch event.Type {
		case dto2.EventAtMessageCreate:
			fmt.Println("www")
			return atMessageHandler(event, event.RawMessage)
		}
	}

	return nil
}

func atMessageHandler(event *dto2.WSPayload, message []byte) error {
	data := &dto2.WSATMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}

	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(event, data)
	}
	return nil
}
