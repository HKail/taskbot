package botclient

import "fmt"

const (
	domain        = "api.sgroup.qq.com"
	sandboxDomain = "sandbox.api.sgroup.qq.com"

	scheme = "https"
)

type uri string

const (
	gatewayURI uri = "/gateway"

	// 发送消息
	messageSend uri = "/channels/{channel_id}/messages"

	// 用户信息
	userMe uri = "/users/@me"
)

func (c *BotClient) getURL(endpoint uri) string {
	d := domain
	if c.sandbox {
		d = sandboxDomain
	}

	return fmt.Sprintf("%s://%s%s", scheme, d, endpoint)
}
