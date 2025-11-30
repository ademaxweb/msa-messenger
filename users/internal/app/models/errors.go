package models

import "errors"

var (
	UserNotFound  = errors.New("user not found")
	AlreadyExists = errors.New("user already exists")
)
