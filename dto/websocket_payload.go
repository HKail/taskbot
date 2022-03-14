package dto

// EventType 事件类型
type EventType string

// WSPayload websocket 消息传输结构
type WSPayload struct {
	OPCode     OPCode      `json:"op"`
	Seq        uint32      `json:"s,omitempty"`
	Type       EventType   `json:"t,omitempty"`
	Data       interface{} `json:"d,omitempty"`
	RawMessage []byte      `json:"-"` // 原始的 message 数据
}

// WSIdentityData 鉴权数据
type WSIdentityData struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"` // array of two integers (shard_id, num_shards)
	Properties struct {
		Os      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"properties,omitempty"`
}
