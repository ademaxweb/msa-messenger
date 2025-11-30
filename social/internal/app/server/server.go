package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	pb "social/pkg/api/social/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Controllers struct {
		Social pb.SocialServiceServer
	}

	Deps struct {
		Controllers
	}

	Config struct {
		Port string

		Interceptors []grpc.UnaryServerInterceptor
	}

	grpcProps struct {
		port     string
		listener net.Listener
		server   *grpc.Server
	}

	Server struct {
		controllers Controllers

		grpc grpcProps
	}
)

func New(cfg Config, deps Deps) *Server {
	srv := &Server{controllers: deps.Controllers}

	var grpcOpts []grpc.ServerOption
	grpcOpts = append(grpcOpts, grpc.ChainUnaryInterceptor(cfg.Interceptors...))

	grpcSrv := grpc.NewServer(grpcOpts...)

	pb.RegisterSocialServiceServer(grpcSrv, srv.controllers.Social)
	reflection.Register(grpcSrv)

	srv.grpc.server = grpcSrv
	srv.grpc.port = cfg.Port

	return srv
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.grpc.port)
	if err != nil {
		return fmt.Errorf("failed to listen port %s: %v", s.grpc.port, err)
	}

	s.grpc.listener = lis

	go func() {
		<-ctx.Done()

		ch := make(chan struct{})

		go func() {
			s.grpc.server.GracefulStop()
			close(ch)
		}()

		select {
		case <-ch:
		case <-time.After(5 * time.Second):
			s.grpc.server.Stop()
		}

		if err := s.grpc.listener.Close(); err != nil {
			log.Printf("failed to close listener: %v", err)
		}
	}()

	err = s.grpc.server.Serve(lis)
	if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
