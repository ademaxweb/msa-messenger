package grpc

import (
	"chat/internal/app/usecases/dto"
	pb "chat/pkg/api/chat/v1"
	"context"
)

func (h *Handler) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	o := dto.SendMessage{
		ChatID:   request.GetChatId(),
		SenderID: 0, // TODO get sender id from headers
		Text:     request.GetText(),
	}

	message, err := h.useCases.SendMessage(o)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		Message: pbMessageFromMessageModel(message),
	}, nil
}
