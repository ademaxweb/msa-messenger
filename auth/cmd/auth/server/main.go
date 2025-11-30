package main

import (
	controllers "auth/internal/app/controllers/grpc"
	"auth/internal/app/server"
	"auth/internal/app/usecases"
	mwsGRPC "auth/internal/middlewares/grpc"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	authService := usecases.NewAuthService()
	authControllers, err := controllers.New(
		controllers.Deps{
			AuthUsecases: authService,
		},
	)
	if err != nil {
		log.Fatalf("failed to initialize auth controllers: %v", err)
	}

	serverCfg := server.Config{
		Port: ":50051",
		Interceptors: []grpc.UnaryServerInterceptor{
			mwsGRPC.ConvertErrorUnaryServerInterceptor(),
		},
	}

	serverDeps := server.Deps{
		Controllers: server.Controllers{
			Auth: authControllers,
		},
	}

	s := server.New(serverCfg, serverDeps)

	err = s.Run(ctx)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
