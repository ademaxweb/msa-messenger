package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"
)

func (s *Handler) DeclineFriendRequest(ctx context.Context, request *pb.DeclineFriendRequestRequest) (*pb.DeclineFriendRequestResponse, error) {
	dto := dtoUpdateRequestStatusFromDeclineFriendRequestRequest(request)

	req, err := s.useCases.UpdateRequestStatus(dto)
	if err != nil {
		return nil, err
	}

	return &pb.DeclineFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
