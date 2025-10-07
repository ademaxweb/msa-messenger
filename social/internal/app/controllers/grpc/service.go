package grpc

import (
	"social/internal/app/usecases"
	pb "social/pkg/api/social/v1"
)

type Handler struct {
	pb.UnimplementedSocialServiceServer

	useCases usecases.Interface
}

func NewHandler(repository usecases.RequestsRepository) *Handler {

	uc := usecases.NewRequestsService(repository)

	return &Handler{
		useCases: uc,
	}

}
