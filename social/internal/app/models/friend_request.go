package models

type FriendRequestStatus int32

const (
	RequestStatusUnknown FriendRequestStatus = iota
	RequestStatusPending
	RequestStatusAccepted
	RequestStatusDeclined
	RequestStatusRemoved
)

type FriendRequest struct {
	// Уникальный идентификатор
	Id uint32
	// Идентификатор пользователя, отправившего заявку
	SenderId uint32
	// Идентификатор пользователя, получившего заявку
	RecipientID uint32
	// Текущий статус заявки
	Status FriendRequestStatus
}
