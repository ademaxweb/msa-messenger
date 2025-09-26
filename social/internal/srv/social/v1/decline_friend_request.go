package v1

import (
	"context"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeclineFriendRequest(ctx context.Context, request *pb.DeclineFriendRequestRequest) (*pb.DeclineFriendRequestResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
