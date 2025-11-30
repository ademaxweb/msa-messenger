package models

import "errors"

var (
	ErrChatNotFound           = errors.New("chat not found")
	ErrCannotChatToYourself   = errors.New("cannot chat to yourself")
	ErrCannotSendEmptyMessage = errors.New("cannot send empty message")
)
