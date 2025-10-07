package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func (rs *RequestsService) ListFriends(o dto.ListFriends) ([]uint32, error) {
	filters := models.FriendRequest{
		SenderId:    o.UserID,
		RecipientID: o.UserID,
		Status:      pb.FriendRequestStatus_STATUS_ACCEPTED,
	}

	reqs, err := rs.repo.GetUserRequests(&filters)
	if err != nil {
		return nil, err
	}

	uids := make([]uint32, len(reqs))

	for _, r := range reqs {
		uids = append(uids, r.Id)
	}

	return uids, nil
}
