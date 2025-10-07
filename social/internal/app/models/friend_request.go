package models

import pb "social/pkg/api/social/v1"

type FriendRequest struct {
	// Уникальный идентификатор
	Id uint32
	// Идентификатор пользователя, отправившего заявку
	SenderId uint32
	// Идентификатор пользователя, получившего заявку
	RecipientID uint32
	// Текущий статус заявки
	Status pb.FriendRequestStatus
}
