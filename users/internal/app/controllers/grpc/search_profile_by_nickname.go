package grpc

import (
	"context"
	pb "users/pkg/api/users/v1"
)

func (s *Service) SearchProfileByNickname(ctx context.Context, request *pb.SearchProfileByNicknameRequest) (*pb.SearchProfileByNicknameResponse, error) {
	dto := dtoSearchProfileByNicknameFromSearchProfileRequest(request)

	users, err := s.useCases.SearchProfileByNickname(dto)
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.UserProfile, len(users))

	for _, u := range users {
		pbUsers = append(pbUsers, modelUserToUserProfile(&u))
	}

	return &pb.SearchProfileByNicknameResponse{
		Results: pbUsers,
		Pagination: &pb.Pagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil
}
