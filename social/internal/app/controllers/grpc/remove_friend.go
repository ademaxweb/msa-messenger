package grpc

import (
	"context"
	"social/internal/app/usecases/dto"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) RemoveFriend(ctx context.Context, request *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	o := dto.RemoveFriend{
		UserID:   0, // TODO get from headers
		FriendID: request.GetUserId(),
	}

	err := s.useCases.RemoveFriend(o)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.RemoveFriendResponse{}, nil
}
