package grpc

import (
	"context"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"
)

func (s *Handler) RemoveFriend(ctx context.Context, request *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	o := dto.RemoveFriend{
		UserID:   0, // TODO get from headers
		FriendID: request.GetUserId(),
	}

	err := s.useCases.RemoveFriend(o)
	if err != nil {
		return nil, err
	}

	return &pb.RemoveFriendResponse{}, nil
}
