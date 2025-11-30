package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"
)

func (h *Handler) GetChat(ctx context.Context, request *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	dto := dtoGetChatFromGetChatRequest(request)

	chat, err := h.useCases.GetChat(dto)
	if err != nil {
		return nil, err
	}

	return &pb.GetChatResponse{
		Chat: pbChatFromChatModel(chat),
	}, nil
}
