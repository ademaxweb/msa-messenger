package v1

import (
	pb "users/pkg/api/users/v1"
)

type Server struct {
	pb.UnimplementedUsersServiceServer
}

func NewServer() *Server {
	return &Server{}
}
