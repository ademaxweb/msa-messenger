package grpc

import (
	"chat/internal/app/usecases"
	pb "chat/pkg/api/chat/v1"
)

type Handler struct {
	pb.UnimplementedChatServiceServer

	useCases usecases.Interface
}

func NewHandler(uc usecases.Interface) *Handler {
	return &Handler{
		useCases: uc,
	}
}
