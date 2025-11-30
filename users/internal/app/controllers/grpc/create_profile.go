package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"
)

func (s *Service) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	dto := dtoCreateProfileFromCreateProfileRequest(request)

	u, err := s.useCases.CreateProfile(dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
