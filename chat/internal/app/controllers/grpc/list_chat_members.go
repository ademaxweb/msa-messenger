package grpc

import (
	pb "chat/pkg/api/chat/v1"
	"context"
)

func (h *Handler) ListChatMembers(ctx context.Context, request *pb.ListChatMembersRequest) (*pb.ListChatMembersResponse, error) {
	dto := dtoListChatMembersFromListChatMembersRequest(request)

	members, err := h.useCases.ListChatMembers(dto)
	if err != nil {
		return nil, err
	}

	return &pb.ListChatMembersResponse{
		UserIds: members,
		// TODO implement pagination
		Pagination: &pb.Pagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil
}
