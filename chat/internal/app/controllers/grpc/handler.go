package grpc

import (
	"chat/internal/app/usecases"
	pb "chat/pkg/api/chat/v1"
)

type Handler struct {
	pb.UnimplementedChatServiceServer

	useCases usecases.Interface
}

func NewHandler(repo usecases.ChatRepository, nameGenerator usecases.NameGenerator) *Handler {
	return &Handler{
		useCases: usecases.NewChatService(repo, nameGenerator),
	}
}
