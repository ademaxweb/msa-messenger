package v1

import (
	pb "chat/pkg/api/chat/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListUserChats(ctx context.Context, request *pb.ListUserChatsRequest) (*pb.ListUserChatsResponse, error) {
	//TODO implement me
	return nil, status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
