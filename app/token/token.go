package token

import "fmt"

type Type string

const (
	TypeBot    = "Bot"
	TypeNormal = "Bearer"
)

type Token struct {
	Scheme      string
	AppID       uint64
	AccessToken string
}

func NewBotToken(appID uint64, accessToken string) *Token {
	return &Token{
		Scheme:      TypeBot,
		AppID:       appID,
		AccessToken: accessToken,
	}
}

func (t *Token) GetString() string {
	return fmt.Sprintf("%v.%s", t.AppID, t.AccessToken)
}
