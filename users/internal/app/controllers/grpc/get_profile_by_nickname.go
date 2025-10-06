package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProfileByNickname(ctx context.Context, request *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	dto := dtoGetProfileByNicknameFromGetProfileByNicknameRequest(request)

	u, err := s.useCases.GetProfileByNickname(dto)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &pb.GetProfileByNicknameResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
