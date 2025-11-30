package grpc

import (
	"context"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func (s *Handler) SendFriendRequest(ctx context.Context, request *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	o := dto.SendFriendRequest{
		SenderID:    0, // TODO get from headers
		RecipientID: request.GetUserId(),
	}

	req, err := s.useCases.SendFriendRequest(o)
	if err != nil {
		return nil, err
	}

	return &pb.SendFriendRequestResponse{
		Request: pbFriendRequestFromFriendRequestModel(req),
	}, nil
}
