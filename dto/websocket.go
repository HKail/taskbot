package dto

import "github.com/hkail/taskbot/token"

// WSAccessPoint websocket 接入点信息
type WSAccessPoint struct {
	URL string `json:"url"`
	// TODO 支持分片
}

// WSSession websocket 链接所需要的会话信息
type WSSession struct {
	ID      string // 鉴权完成后由 QQ 服务器下发
	LastSeq uint32 // 消息序列号, 保存用于发送心跳时使用
	URL     string
	Token   token.Token
	Intent  Intent
	// TODO 分片相关属性
}
