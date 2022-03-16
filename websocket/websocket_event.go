package websocket

import (
	"fmt"

	"github.com/hkail/taskbot/dto"
)

func parseAndHandleEvent(event *dto.WSPayload) error {
	fmt.Println(event)

	if event.OPCode == dto.OPCodeDispatch {
		switch event.Type {
		case dto.EventAtMessageCreate:
			fmt.Println("www")
			return atMessageHandler(event, event.RawMessage)
		}
	}

	return nil
}

func atMessageHandler(event *dto.WSPayload, message []byte) error {
	data := &dto.WSATMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}

	if DefaultHandlers.ATMessage != nil {
		return DefaultHandlers.ATMessage(event, data)
	}
	return nil
}
