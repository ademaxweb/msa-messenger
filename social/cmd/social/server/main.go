package main

import (
	"log"
	"net"
	requestsRepository "social/internal/app/repositories/requests"

	handler "social/internal/app/controllers/grpc"
	pb "social/pkg/api/social/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	requestsRepo := requestsRepository.NewRepository()

	impl := handler.NewHandler(requestsRepo)

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterSocialServiceServer(server, impl)

	reflection.Register(server)

	log.Printf("server listening at %v", ls.Addr())
	if err := server.Serve(ls); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
