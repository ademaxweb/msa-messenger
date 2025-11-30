package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"
)

func (s *Handler) AcceptFriendRequest(ctx context.Context, request *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	dto := dtoUpdateRequestStatusFromAcceptFriendRequestRequest(request)

	req, err := s.useCases.UpdateRequestStatus(dto)
	if err != nil {
		return nil, err
	}

	return &pb.AcceptFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
