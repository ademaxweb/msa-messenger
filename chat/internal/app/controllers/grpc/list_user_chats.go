package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) ListUserChats(ctx context.Context, request *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	dto := dtoListUserChatsFromListUserChatsRequest(request)

	chats, err := h.useCases.ListUserChats(dto)
	if err != nil {
		// TODO implement errors handling
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
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
