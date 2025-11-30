package main

import (
	controllers "chat/internal/app/controllers/grpc"
	"chat/internal/app/repositories/chat"
	"chat/internal/app/server"
	"chat/internal/app/usecases"
	"chat/pkg/namegen"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nameGenerator := namegen.NewGenerator()
	chatRepository := chat.NewRepository()
	chatService := usecases.NewChatService(chatRepository, nameGenerator)
	chatControllers := controllers.NewHandler(chatService)

	srvCfg := server.Config{
		Port:         ":50051",
		Interceptors: []grpc.UnaryServerInterceptor{},
	}

	srvDeps := server.Deps{
		Controllers: server.Controllers{
			Chat: chatControllers,
		},
	}

	srv := server.New(srvCfg, srvDeps)

	err := srv.Run(ctx)
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
