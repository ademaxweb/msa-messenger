package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) ListMessages(ctx context.Context, request *pb.ListMessagesRequest) (*pb.ListMessagesResponse, error) {
	dto := dtoListMessagesFromListMessagesRequest(request)

	messages, err := h.useCases.ListMessages(dto)
	if err != nil {
		// TODO implement errors handling
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}

	pbMessages := make([]*pb.Message, len(messages))
	for _, m := range messages {
		pbMessages = append(pbMessages, pbMessageFromMessageModel(&m))
	}

	return &pb.ListMessagesResponse{
		Messages: pbMessages,
		// TODO implement pagination
		Pagination: &pb.Pagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil
}
