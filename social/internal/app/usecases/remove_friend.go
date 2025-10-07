package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
)

func (rs *RequestsService) RemoveFriend(o dto.RemoveFriend) error {
	r := models.FriendRequest{
		SenderId:    o.UserID,
		RecipientID: o.FriendID,
	}

	_, err := rs.repo.UpdateRequestByUsers(&r)

	return err
}
