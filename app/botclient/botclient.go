package botclient

import (
	"context"
	"time"

	"github.com/hkail/taskbot/app/dto"
	"github.com/hkail/taskbot/app/token"

	"github.com/go-resty/resty/v2"
)

type BotClient struct {
	token       *token.Token
	sandbox     bool
	restyClient *resty.Client
}

func NewBotClient(isSandBox bool, token *token.Token) *BotClient {
	return &BotClient{
		token:   token,
		sandbox: isSandBox,
		restyClient: resty.New().
			SetTimeout(time.Second * 3).
			SetAuthScheme(token.Scheme).
			SetAuthToken(token.GetString()),
	}
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

// MessageSend 发送消息
func (c *BotClient) MessageSend(ctx context.Context, channelID string, msg *dto.PostChannelMessage) (*dto.Message, error) {
	resp, err := c.request(ctx).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(c.getURL(messageSend))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.Message), nil
}

// Me 拉取当前用户的信息
func (c *BotClient) Me(ctx context.Context) (*dto.User, error) {
	resp, err := c.request(ctx).
		SetResult(dto.User{}).
		Get(c.getURL(userMe))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*dto.User), nil
}
