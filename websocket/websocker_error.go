package websocket

import "fmt"

const (
	errCodeNeedReconnect = 7000 + iota
	errCodeInvalidSession
)

var (
	ErrNeedReconnect  = NewWSError(errCodeNeedReconnect, "need reconnect")
	ErrInvalidSession = NewWSError(errCodeInvalidSession, "invalid session")
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
