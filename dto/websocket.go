package dto

// WSAccessPoint websocket 接入点信息
type WSAccessPoint struct {
	URL string `json:"url"`
	// TODO 支持分片
}
