package v1

import (
	pb "auth/pkg/api/auth/v1"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
}

func NewServer() *Server {
	return &Server{}
}
