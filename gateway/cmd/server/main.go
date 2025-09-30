package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ademaxweb/msa-messenger/auth/pkg/api/auth/v1"
	"github.com/ademaxweb/msa-messenger/chat/pkg/api/chat/v1"
	"github.com/ademaxweb/msa-messenger/social/pkg/api/social/v1"
	"github.com/ademaxweb/msa-messenger/users/pkg/api/users/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	cfg := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth-service:8080", cfg); err != nil {
		log.Fatalf("RegisterUsersServiceHandlerFromEndpoint error: %v", err)
	}

	if err := chat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat-service:8080", cfg); err != nil {
		log.Fatalf("RegisterUsersServiceHandlerFromEndpoint error: %v", err)
	}

	if err := social.RegisterSocialServiceHandlerFromEndpoint(ctx, mux, "social-service:8080", cfg); err != nil {
		log.Fatalf("RegisterUsersServiceHandlerFromEndpoint error: %v", err)
	}

	if err := users.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, "users-service:8080", cfg); err != nil {
		log.Fatalf("RegisterUsersServiceHandlerFromEndpoint error: %v", err)
	}

	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe() error: %v", err)
	}
}
