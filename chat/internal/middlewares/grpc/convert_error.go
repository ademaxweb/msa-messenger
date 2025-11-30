package grpc

import (
	"chat/internal/app/models"
	"context"
	"errors"
	"log"

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
	) (resp any, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			// Transport err
			if _, ok := status.FromError(err); ok {
				return
			}

			// Business err
			switch {
			case errors.Is(err, models.ErrChatNotFound):
				return nil, status.Error(codes.NotFound, err.Error())
			case errors.Is(err, models.ErrCannotChatToYourself):
				return nil, status.Error(codes.Unavailable, err.Error())
			case errors.Is(err, models.ErrCannotSendEmptyMessage):
				return nil, status.Error(codes.InvalidArgument, err.Error())
			default:
				log.Printf("Request returned unknown error: %v\n", err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
