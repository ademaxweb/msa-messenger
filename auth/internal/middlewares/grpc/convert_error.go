package grpc

import (
	"context"
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
			if _, ok := status.FromError(err); ok {
				return
			}

			switch {
			default:
				log.Printf("Request returned unknown error: %v\n", err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}

		return
	}
}
