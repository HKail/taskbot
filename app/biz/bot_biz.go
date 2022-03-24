package biz

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hkail/taskbot/app/dto"
)

// ATMessageEventHandler 机器人艾特事件处理
func (biz *Biz) ATMessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {
	ctx := context.Background()

	if !strings.Contains(data.Content, "/打卡") {
		content := fmt.Sprintf("%s 抱歉，目前仅支持指令[/打卡]～", dto.MentionUser(data.Author.ID))
		return biz.sendMessageToUser(data, content)
	}

	guildID, err := strconv.ParseUint(data.GuildID, 10, 64)
	if err != nil {
		log.Printf("ATMessageEventHandler parse guild_id has err=%v", err)
		return err
	}

	userID, err := strconv.ParseUint(data.Author.ID, 10, 64)
	if err != nil {
		log.Printf("ATMessageEventHandler parse user_id has err=%v", err)
		return err
	}

	firstCheckin, err := biz.UserCheckin(ctx, time.Now(), guildID, userID)
	if err != nil {
		log.Printf("ATMessageEventHandler UserCheckin has err=%v", err)
		return err
	}

	if !firstCheckin {
		content := fmt.Sprintf("%s 您今日已打过卡～", dto.MentionUser(data.Author.ID))
		return biz.sendMessageToUser(data, content)
	}

	contCheckinDays, err := biz.GetUserContCheckinDays(ctx, guildID, userID)
	if err != nil {
		log.Printf("ATMessageEventHandler GetUserContCheckinDays has err=%v", err)
		return err
	}

	content := fmt.Sprintf("%s 打卡成功，您已连续打卡%d天！", dto.MentionUser(data.Author.ID), contCheckinDays)
	return biz.sendMessageToUser(data, content)
}

func (biz *Biz) sendMessageToUser(data *dto.WSATMessageData, content string) error {
	_, err := biz.botClient.MessageSend(context.Background(), data.ChannelID, &dto.PostChannelMessage{
		Content: content,
		MsgID:   data.ID,
	})
	if err != nil {
		log.Printf("sendMessageToUser has err=%v", err)
		return err
	}

	return nil
}
