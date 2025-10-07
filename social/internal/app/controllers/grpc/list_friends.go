package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) ListFriends(ctx context.Context, request *pb.ListFriendsRequest) (*pb.ListFriendsResponse, error) {
	dto := dtoListFriendsFromListFriendsRequest(request)

	uids, err := s.useCases.ListFriends(dto)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
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
