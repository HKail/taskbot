package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hkail/taskbot/botclient"
	"github.com/hkail/taskbot/dto"
	"github.com/hkail/taskbot/token"
	"github.com/hkail/taskbot/websocket"
)

func main() {
	fmt.Println("hello world")
	botToken := token.NewBotToken(0, "")
	botClient := botclient.NewSandboxBotClient(botToken)

	wsAP, err := botClient.GetWSAccessPoint(context.Background())
	if err != nil {
		panic(err)
	}

	//user, err := botClient.Me(context.Background())
	//if err != nil {
	//	panic(err)
	//}

	var atMessage websocket.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		fmt.Println(event.Type, event.OPCode)
		fmt.Printf("%#v\n", data)

		msg, err := botClient.MessageSend(context.Background(), data.ChannelID, &dto.PostChannelMessage{
			Content: fmt.Sprintf("你也好：%s", dto.MentionUser(data.Author.ID)),
			MsgID:   data.ID,
		})
		if err != nil {
			log.Printf("err=%v", err)
		}

		fmt.Printf("%#v\n", msg)

		return nil
	}
	intent := websocket.RegisterHandlers(atMessage)

	if err = websocket.NewWSManager().Start(wsAP, botToken, intent); err != nil {
		panic(err)
	}
}
