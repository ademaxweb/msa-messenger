package dto

import (
	"social/internal/app/models"
)

type UpdateRequestStatus struct {
	RequestId uint32
	Status    models.FriendRequestStatus
}
