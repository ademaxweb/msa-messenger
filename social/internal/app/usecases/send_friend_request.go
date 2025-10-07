package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func (rs *RequestsService) SendFriendRequest(o dto.SendFriendRequest) (*models.FriendRequest, error) {
	r := models.FriendRequest{
		SenderId:    o.SenderID,
		RecipientID: o.RecipientID,
		Status:      pb.FriendRequestStatus_STATUS_PENDING,
	}

	reqID, err := rs.repo.CreateRequest(&r)
	if err != nil {
		return nil, err
	}

	r.Id = reqID

	return &r, nil
}
