package main

import (
	"chat/internal/app/repositories/chat"
	pb "chat/pkg/api/chat/v1"
	"chat/pkg/namegen"
	"log"
	"net"

	handler "chat/internal/app/controllers/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	chatRepo := chat.NewRepository()

	nameGenerator := namegen.NewGenerator()

	impl := handler.NewHandler(chatRepo, nameGenerator)

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
