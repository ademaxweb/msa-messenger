package v1

import (
	pb "chat/pkg/api/chat/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) StreamMessages(request *pb.StreamMessagesRequest, g grpc.ServerStreamingServer[pb.StreamMessagesResponse]) error {
	//TODO implement me
	return status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
