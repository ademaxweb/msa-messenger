package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"
)

func (s *Service) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	dto := dtoUpdateProfileFromUpdateProfileRequest(request)

	u, err := s.useCases.UpdateProfile(dto)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
