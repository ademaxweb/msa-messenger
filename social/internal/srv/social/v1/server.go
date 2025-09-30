package v1

import (
	pb "social/pkg/api/social/v1"
)

type Server struct {
	pb.UnimplementedSocialServiceServer
}

func NewServer() *Server {
	return &Server{}
}
