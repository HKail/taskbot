package dto

// PostChannelMessage 发送消息结构体定义
type PostChannelMessage struct {
	Content          string            `json:"content,omitempty"`
	Embed            *Embed            `json:"embed,omitempty"`
	Ark              *Ark              `json:"ark,omitempty"`
	MessageReference *MessageReference `json:"message_reference,omitempty"`
	Image            string            `json:"image,omitempty"`
	MsgID            string            `json:"msg_id,omitempty"` // 要回复的消息id，为空是主动消息，公域机器人会异步审核，不为空是被动消息，公域机器人会校验语料
}
