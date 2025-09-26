package main

import (
	authsrv "auth/internal/srv/users/v1"
	pb "auth/pkg/api/auth/v1"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	impl := authsrv.NewServer()

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterAuthServiceServer(srv, impl)

	reflection.Register(srv)

	log.Printf("Listening at %v", ls.Addr())
	if err := srv.Serve(ls); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
