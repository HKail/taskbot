package botclient

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/hkail/taskbot/dto"
	"github.com/hkail/taskbot/token"
)

type BotClient struct {
	token       *token.Token
	sandbox     bool
	restyClient *resty.Client
}

func newBotClient(isSandBox bool, token *token.Token) *BotClient {
	return &BotClient{
		token:   token,
		sandbox: isSandBox,
		restyClient: resty.New().
			SetAuthScheme(token.Scheme).
			SetAuthToken(token.GetString()),
	}
}

func NewBotClient(token *token.Token) *BotClient {
	return newBotClient(false, token)
}

func NewSandboxBotClient(token *token.Token) *BotClient {
	return newBotClient(true, token)
}

func (c *BotClient) request(ctx context.Context) *resty.Request {
	return c.restyClient.R().SetContext(ctx)
}

func (c *BotClient) GetWSAccessPoint(ctx context.Context) (*dto.WSAccessPoint, error) {
	resp, err := c.request(ctx).
		SetResult(dto.WSAccessPoint{}).
		Get(c.getURL(gatewayURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.WSAccessPoint), nil
}
