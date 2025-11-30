package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"
)

func (h *Handler) ListUserChats(ctx context.Context, request *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	dto := dtoListUserChatsFromListUserChatsRequest(request)

	chats, err := h.useCases.ListUserChats(dto)
	if err != nil {
		return nil, err
	}

	pbChats := make([]*pb.Chat, 0, len(chats))
	for _, c := range chats {
		pbChats = append(pbChats, pbChatFromChatModel(&c))
	}

	return &pb.ListUserChatsResponse{
		Chats: pbChats,
		Pagination: &pb.Pagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil

}
