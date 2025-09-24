package v1

import (
	pb "auth/pkg/api/auth/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}

func (s Server) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
