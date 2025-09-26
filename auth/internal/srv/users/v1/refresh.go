package v1

import (
	pb "auth/pkg/api/auth/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
