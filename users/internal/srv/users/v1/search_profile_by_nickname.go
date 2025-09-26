package v1

import (
	"context"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchProfileByNickname(ctx context.Context, request *pb.SearchProfileByNicknameRequest) (*pb.SearchProfileByNicknameResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
