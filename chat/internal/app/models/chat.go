package models

type Chat struct {
	// Id уникальный идентификатор чата
	Id uint32
	// Name название чата
	Name string
	// Description описание чата
	Description string
	// AvatarURL URL аватарки чата
	AvatarURL string
}
