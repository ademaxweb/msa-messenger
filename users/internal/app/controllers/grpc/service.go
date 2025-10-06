package grpc

import (
	"users/internal/app/usecases"
	pb "users/pkg/api/users/v1"
)

type Service struct {
	pb.UnimplementedUsersServiceServer
	useCases usecases.Interface
}

func NewHandler(useCases usecases.Interface) *Service {
	return &Service{
		useCases: useCases,
	}
}
