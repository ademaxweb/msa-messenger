package grpc

import (
	"context"
	"errors"
	"log"
	"users/internal/app/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			// Если ошибка уже транспортная
			if _, ok := status.FromError(err); ok {
				return resp, err
			}

			// Если ошибка бизнесовая
			switch {
			case errors.Is(err, models.UserNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, models.AlreadyExists):
				return nil, status.Error(codes.AlreadyExists, err.Error())
			default:
				log.Printf("Requests returned unknown errpor: %v", err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
