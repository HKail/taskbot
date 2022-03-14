package main

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"

	"github.com/hkail/taskbot/botclient"
	"github.com/hkail/taskbot/token"
)

func main() {
	fmt.Println("hello world")
	botToken := token.NewBotToken(0, "")
	botClient := botclient.NewSandboxBotClient(botToken)

	wsAP, err := botClient.GetWSAccessPoint(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(wsAP)

	conn, _, err := websocket.DefaultDialer.Dial("wss://sandbox.api.sgroup.qq.com/websocket", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(conn)
}
