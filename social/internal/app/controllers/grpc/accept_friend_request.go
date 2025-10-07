package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) AcceptFriendRequest(ctx context.Context, request *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	dto := dtoUpdateRequestStatusFromAcceptFriendRequestRequest(request)

	req, err := s.useCases.UpdateRequestStatus(dto)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.AcceptFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
