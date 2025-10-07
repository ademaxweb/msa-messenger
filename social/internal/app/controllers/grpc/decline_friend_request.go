package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) DeclineFriendRequest(ctx context.Context, request *pb.DeclineFriendRequestRequest) (*pb.DeclineFriendRequestResponse, error) {
	dto := dtoUpdateRequestStatusFromDeclineFriendRequestRequest(request)

	req, err := s.useCases.UpdateRequestStatus(dto)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.DeclineFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
