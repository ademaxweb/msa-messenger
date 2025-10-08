package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) GetChat(ctx context.Context, request *pb.GetChatRequest) (*pb.GetChatResponse, error) {
	dto := dtoGetChatFromGetChatRequest(request)

	chat, err := h.useCases.GetChat(dto)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.GetChatResponse{
		Chat: pbChatFromChatModel(chat),
	}, nil
}
