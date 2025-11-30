package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"
)

func (s *Service) GetProfileByNickname(ctx context.Context, request *pb.GetProfileByNicknameRequest) (*pb.GetProfileByNicknameResponse, error) {
	dto := dtoGetProfileByNicknameFromGetProfileByNicknameRequest(request)

	u, err := s.useCases.GetProfileByNickname(dto)
	if err != nil {
		return nil, err
	}

	return &pb.GetProfileByNicknameResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
