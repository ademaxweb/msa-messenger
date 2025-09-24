package main

import (
	chatsrv "chat/internal/srv/chat/v1"
	pb "chat/pkg/api/chat/v1"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	impl := chatsrv.NewServer()

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	pb.RegisterChatServiceServer(srv, impl)

	reflection.Register(srv)

	if err := srv.Serve(ls); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
