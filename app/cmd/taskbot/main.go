package main

import (
	"context"
	"flag"

	"github.com/hkail/taskbot/app/biz"
	"github.com/hkail/taskbot/app/conf"

	"github.com/hkail/taskbot/app/botclient"
	"github.com/hkail/taskbot/app/token"
	"github.com/hkail/taskbot/app/websocket"
)

func main() {
	configPath := "./conf/local/config.yml"
	flag.StringVar(&configPath, "config", configPath, "config path")
	flag.Parse()

	appConf, err := conf.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	defaultBiz, err := biz.NewBiz(appConf)
	if err != nil {
		panic(err)
	}

	botToken := token.NewBotToken(101996085, "EV4ju1yahUBgXe7xOjGqzv45QpSUONVK")
	botClient := botclient.NewBotClient(appConf.Bot.IsSandbox, botToken)

	wsAP, err := botClient.GetWSAccessPoint(context.Background())
	if err != nil {
		panic(err)
	}

	var atMessage websocket.ATMessageEventHandler = defaultBiz.ATMessageEventHandler
	intent := websocket.RegisterHandlers(atMessage)

	if err = websocket.NewWSManager().Start(wsAP, botToken, intent); err != nil {
		panic(err)
	}
}
