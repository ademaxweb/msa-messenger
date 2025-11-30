package grpc

import (
	"chat/internal/app/usecases/dto"
	pb "chat/pkg/api/chat/v1"
	"context"
)

func (h *Handler) CreateDirectChat(ctx context.Context, request *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {
	o := dto.CreateDirectChat{
		SenderID:    0,
		RecipientID: request.GetParticipantId(),
	}

	chat, err := h.useCases.CreateDirectChat(o)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDirectChatResponse{
		ChatId: chat.ID,
	}, nil
}
