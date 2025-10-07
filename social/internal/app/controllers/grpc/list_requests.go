package grpc

import (
	"context"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Handler) ListRequests(ctx context.Context, request *pb.ListRequestsRequest) (*pb.ListRequestsResponse, error) {
	dto := dtoListRequestsFromListRequestsRequest(request)

	requests, err := s.useCases.ListRequests(dto)
	if err != nil {
		return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
	}
	
	pbRequests := make([]*pb.FriendRequest, len(requests))

	for _, r := range requests {
		pbRequests = append(pbRequests, pbFriendRequestFromFriendRequestModel(&r))
	}

	return &pb.ListRequestsResponse{
		Requests: pbRequests,
		Pagination: &pb.CursorPagination{
			Limit:      0,
			Cursor:     nil,
			NextCursor: nil,
		},
	}, nil

}
