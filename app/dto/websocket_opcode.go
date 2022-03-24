package dto

// OPCode websocket 操作码
type OPCode int

// 具体参考: https://bot.q.qq.com/wiki/develop/api/gateway/opcode.html
const (
	OPCodeDispatch       = 0  // 服务端进行消息推送
	OPCodeHeartbeat      = 1  // 客户端或服务端发送心跳
	OPCodeIdentify       = 2  // 客户端发送鉴权
	OPCodeResume         = 6  // 客户端恢复连接
	OPCodeReconnect      = 7  // 服务端通知客户端重新连接
	OPCodeInvalidSession = 9  // 当identify或resume的时候，如果参数有错，服务端会返回该消息
	OPCodeHello          = 10 // 当客户端与网关建立ws连接之后，网关下发的第一条消息
	OPCodeHeartbeatACK   = 11 // 当发送心跳成功之后，就会收到该消息
)
