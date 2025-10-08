package models

import "errors"

var (
	ChatNotFound           = errors.New("chat not found")
	CannotChatToYourself   = errors.New("cannot chat to yourself")
	CannotSendEmptyMessage = errors.New("cannot send empty message")
)
