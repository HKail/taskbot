package biz

import (
	"context"
	"fmt"
	"log"

	"github.com/hkail/taskbot/app/dto"
)

// ATMessageEventHandler 机器人艾特事件处理
func (biz *Biz) ATMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	fmt.Println(event.Type, event.OPCode)
	fmt.Printf("%#v\n", data)
	fmt.Println("user is is")
	fmt.Printf("%#v\n", data.Author)

	msg, err := biz.botClient.MessageSend(context.Background(), data.ChannelID, &dto.PostChannelMessage{
		Content: fmt.Sprintf("你也好：%s", dto.MentionUser(data.Author.ID)),
		MsgID:   data.ID,
	})
	if err != nil {
		log.Printf("err=%v", err)
	}

	fmt.Printf("%#v\n", msg)

	return nil
}
