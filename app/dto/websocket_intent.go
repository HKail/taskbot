package dto

// Intent 事件 intents
type Intent int

const (
	// IntentGuilds 包含
	// - GUILD_CREATE -- 当机器人加入新 guild 时
	// - GUILD_UPDATE -- 当 guild 资料发生变更时
	// - GUILD_DELETE -- 当机器人退出 guild 时
	// - CHANNEL_CREATE -- 当 channel 被创建时
	// - CHANNEL_UPDATE -- 当 channel 被更新时
	// - CHANNEL_DELETE -- 当 channel 被删除时
	IntentGuilds = 1 << 0

	// IntentGuildMembers 包含
	// - GUILD_MEMBER_ADD -- 当成员加入时
	// - GUILD_MEMBER_UPDATE -- 当成员资料变更时
	// - GUILD_MEMBER_REMOVE -- 当成员被移除时
	IntentGuildMembers = 1 << 1

	// IntentGuildMessages 包含
	// - MESSAGE_CREATE -- 发送消息事件，代表频道内的全部消息，而不只是 at 机器人的消息。内容与 AT_MESSAGE_CREATE 相同
	IntentGuildMessages = 1 << 9

	// IntentGuildMessageReactions 包含
	// - MESSAGE_REACTION_ADD -- 为消息添加表情表态
	// - MESSAGE_REACTION_REMOVE -- 为消息删除表情表态
	IntentGuildMessageReactions = 1 << 10

	// IntentDirectMessage 包含
	// - DIRECT_MESSAGE_CREATE -- 当收到用户发给机器人的私信消息时
	IntentDirectMessage = 1 << 12

	// IntentMessageAudit 包含
	// - MESSAGE_AUDIT_PASS -- 消息审核通过
	// - MESSAGE_AUDIT_REJECT -- 消息审核不通过
	IntentMessageAudit = 1 << 27

	// IntentAudioAction 包含
	// - AUDIO_START -- 音频开始播放时
	// - AUDIO_FINISH -- 音频播放结束时
	// - AUDIO_ON_MIC -- 上麦时
	// - AUDIO_OFF_MIC -- 下麦时
	IntentAudioAction = 1 << 29

	// IntentAtMessages 包含
	// - AT_MESSAGE_CREATE -- 当收到@机器人的消息时
	IntentAtMessages = 1 << 30
)
