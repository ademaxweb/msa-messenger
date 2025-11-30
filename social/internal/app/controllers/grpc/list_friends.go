package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"
)

func (s *Handler) ListFriends(ctx context.Context, request *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	dto := dtoListFriendsFromListFriendsRequest(request)

	uids, err := s.useCases.ListFriends(dto)
	if err != nil {
		return nil, err
	}

	return &pb.ListFriendsResponse{
		FriendUserIds: uids,
		Pagination: &pb.CursorPagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil
}
