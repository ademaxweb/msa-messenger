package grpc

import (
	"chat/internal/app/usecases/dto"
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CreateDirectChat(ctx context.Context, request *pb.CreateDirectChatRequest) (*pb.CreateDirectChatResponse, error) {
	o := dto.CreateDirectChat{
		SenderID:    0,
		RecipientID: request.GetParticipantId(),
	}

	chat, err := h.useCases.CreateDirectChat(o)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.CreateDirectChatResponse{
		ChatId: chat.Id,
	}, nil
}
