package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProfile(ctx context.Context, request *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	dto := dtoCreateProfileFromCreateProfileRequest(request)

	u, err := s.useCases.CreateProfile(dto)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &pb.CreateProfileResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
