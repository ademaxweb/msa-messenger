package grpc

import (
	"chat/internal/app/usecases/dto"
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	o := dto.SendMessage{
		ChatId:   request.GetChatId(),
		SenderId: 0, // TODO get sender id from headers
		Text:     request.GetText(),
	}

	message, err := h.useCases.SendMessage(o)
	if err != nil {
		// TODO implement errors handling
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	return &pb.SendMessageResponse{
		Message: pbMessageFromMessageModel(message),
	}, nil
}
