package grpc

import (
	"context"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) SendFriendRequest(ctx context.Context, request *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	o := dto.SendFriendRequest{
		SenderID:    0, // TODO get from headers
		RecipientID: request.GetUserId(),
	}

	req, err := s.useCases.SendFriendRequest(o)
	if err != nil {
		return nil, status.New(codes.Unknown, codes.Unimplemented.String()).Err()
	}

	return &pb.SendFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
