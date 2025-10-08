package grpc

import (
	"chat/internal/app/usecases"
	pb "chat/pkg/api/chat/v1"
)

type Handler struct {
	pb.UnimplementedChatServiceServer

	useCases usecases.Interface
}

func NewHandler(repo usecases.ChatRepository) *Handler {
	return &Handler{
		useCases: usecases.NewChatService(repo),
	}
}
