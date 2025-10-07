package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
)

func (rs *RequestsService) UpdateRequestStatus(o dto.UpdateRequestStatus) (*models.FriendRequest, error) {
	r := models.FriendRequest{
		Id:     o.RequestId,
		Status: o.Status,
	}

	updReq, err := rs.repo.UpdateRequest(&r)
	if err != nil {
		return nil, err
	}

	return updReq, nil
}
