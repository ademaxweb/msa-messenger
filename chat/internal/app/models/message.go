package models

type Message struct {
	// Id идентификатор сообщения
	Id uint32
	// ChatID ID чата, в котором было отправлено сообщение
	ChatID uint32
	// UserID ID пользователя, который отправил сообщение
	UserID uint32
	// Text текст сообщения
	Text string
}
