package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func (rs *RequestsService) ListRequests(o dto.ListRequests) ([]models.FriendRequest, error) {
	filters := models.FriendRequest{
		SenderId:    o.UserID,
		RecipientID: o.UserID,
		Status:      pb.FriendRequestStatus_STATUS_PENDING,
	}

	reqs, err := rs.repo.GetUserRequests(&filters)
	if err != nil {
		return nil, err
	}

	return reqs, nil
}
