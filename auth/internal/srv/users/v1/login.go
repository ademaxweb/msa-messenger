package v1

import (
	pb "auth/pkg/api/auth/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
