package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProfileByID(ctx context.Context, request *pb.GetProfileByIDRequest) (*pb.GetProfileByIDResponse, error) {
	dto := dtoGetProfileByIDFromGetProfileByIDRequest(request)

	u, err := s.useCases.GetProfileByID(dto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetProfileByIDResponse{
		UserProfile: modelUserToUserProfile(u),
	}, nil
}
