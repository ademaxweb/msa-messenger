package main

import (
	"context"
	"log"
	controllers "social/internal/app/controllers/grpc"
	"social/internal/app/repositories/requests"
	"social/internal/app/server"
	"social/internal/app/usecases"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reqRepository := requests.NewRepository()
	socialService := usecases.NewRequestsService(reqRepository)
	socialControllers := controllers.NewHandler(socialService)

	srvCfg := server.Config{
		Port:         ":50051",
		Interceptors: []grpc.UnaryServerInterceptor{},
	}

	srvDeps := server.Deps{
		Controllers: server.Controllers{
			Social: socialControllers,
		},
	}

	srv := server.New(srvCfg, srvDeps)

	err := srv.Run(ctx)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}

}
