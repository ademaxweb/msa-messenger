package usecases

import (
	"chat/internal/app/usecases/dto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cs *ChatService) StreamMessages(o dto.StreamMessages) error {
	return status.New(codes.Unimplemented, codes.Unimplemented.String()).Err()
}
