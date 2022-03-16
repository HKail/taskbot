package websocket

import "fmt"

// 内部使用的错误码
const (
	// errCodeConnNeedReconnect 需要重连
	errCodeConnNeedReconnect = 7000 + iota
	// errCodeConnNeedReIdentify 需要重新鉴权
	errCodeConnNeedReIdentify
	// errCodeConnNeedPanic 无法恢复的链接
	errCodeConnNeedPanic
)

// API websocket 错误码
// 参考: https://bot.q.qq.com/wiki/develop/api/gateway/error/error.html
const (
	// 无法重连和鉴权
	errCodeBotRemoved = 4914
	errCodeBotBanned  = 4915

	// 可以直接重连
	errCodeSendMessageTooFast = 4008
	errCodeSessionTimeout     = 4009
)

var (
	ErrNeedReconnect  = NewWSError(errCodeConnNeedReconnect, "need reconnect")
	ErrInvalidSession = NewWSError(errCodeConnNeedReIdentify, "invalid session")
)

type WSError struct {
	code int
	text string
}

func NewWSError(code int, text string) error {
	err := &WSError{
		code: code,
		text: text,
	}

	return err
}

func (e WSError) Error() string {
	return fmt.Sprintf("code: %v, text: %v", e.code, e.text)
}

func (e WSError) Code() int {
	return e.code
}

func (e WSError) Text() string {
	return e.text
}

func IsNeedReconnectError(err error) bool {
	if e, ok := err.(*WSError); ok {
		return e.code == errCodeConnNeedReconnect
	}

	return false
}

func IsNeedReIdentifyError(err error) bool {
	if e, ok := err.(*WSError); ok {
		return e.code == errCodeConnNeedReIdentify
	}

	return false
}

func IsNeedPanicError(err error) bool {
	if e, ok := err.(*WSError); ok {
		return e.code == errCodeConnNeedPanic
	}

	return false
}
