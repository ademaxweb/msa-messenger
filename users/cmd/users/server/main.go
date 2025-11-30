package main

import (
	"context"
	"log"
	handler "users/internal/app/controllers/grpc"
	repo "users/internal/app/repositories/users"
	"users/internal/app/server"
	"users/internal/app/usecases"
	middlewaresGRPC "users/internal/middlewares/grpc"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	usrRepo := repo.NewRepository()
	usrService := usecases.NewUsersService(usrRepo)
	usrControllers := handler.NewHandler(usrService)

	srvCfg := server.Config{
		Port: ":50051",
		Interceptors: []grpc.UnaryServerInterceptor{
			middlewaresGRPC.ConvertErrorUnaryServerInterceptor(),
		},
	}

	srvDeps := server.Deps{
		Controllers: server.Controllers{
			Users: usrControllers,
		},
	}

	s := server.New(srvCfg, srvDeps)

	err := s.Run(ctx)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
