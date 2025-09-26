package v1

import (
	pb "chat/pkg/api/chat/v1"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() *Server {
	return &Server{}
}
