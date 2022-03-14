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
)

func (c *BotClient) getURL(endpoint uri) string {
	d := domain
	if c.sandbox {
		d = sandboxDomain
	}

	return fmt.Sprintf("%s://%s%s", scheme, d, endpoint)
}
