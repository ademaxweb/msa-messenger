package models

type Chat struct {
	// ID уникальный идентификатор чата
	ID uint32
	// Name название чата
	Name string
	// Description описание чата
	Description string
	// AvatarURL URL аватарки чата
	AvatarURL string
}
