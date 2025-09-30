package main

import (
	"log"
	"net"
	usrsrv "users/internal/srv/users/v1"
	pb "users/pkg/api/users/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	impl := usrsrv.NewServer()

	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUsersServiceServer(server, impl)

	reflection.Register(server)

	log.Printf("server listening at %v", ls.Addr())
	if err := server.Serve(ls); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
