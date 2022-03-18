package main

import (
	"context"
	"fmt"
	"log"

	botclient2 "github.com/hkail/taskbot/app/botclient"
	dto2 "github.com/hkail/taskbot/app/dto"
	token2 "github.com/hkail/taskbot/app/token"
	websocket2 "github.com/hkail/taskbot/app/websocket"
)

func main() {
	fmt.Println("hello world")
	botToken := token2.NewBotToken(101996085, "EV4ju1yahUBgXe7xOjGqzv45QpSUONVK")
	botClient := botclient2.NewSandboxBotClient(botToken)

	wsAP, err := botClient.GetWSAccessPoint(context.Background())
	if err != nil {
		panic(err)
	}

	//user, err := botClient.Me(context.Background())
	//if err != nil {
	//	panic(err)
	//}

	var atMessage websocket2.ATMessageEventHandler = func(event *dto2.WSPayload, data *dto2.WSATMessageData) error {
		fmt.Println(event.Type, event.OPCode)
		fmt.Printf("%#v\n", data)
		fmt.Println("user is is")
		fmt.Printf("%#v\n", data.Author)

		msg, err := botClient.MessageSend(context.Background(), data.ChannelID, &dto2.PostChannelMessage{
			Content: fmt.Sprintf("你也好：%s", dto2.MentionUser(data.Author.ID)),
			MsgID:   data.ID,
		})
		if err != nil {
			log.Printf("err=%v", err)
		}

		fmt.Printf("%#v\n", msg)

		return nil
	}
	intent := websocket2.RegisterHandlers(atMessage)

	if err = websocket2.NewWSManager().Start(wsAP, botToken, intent); err != nil {
		panic(err)
	}
}
