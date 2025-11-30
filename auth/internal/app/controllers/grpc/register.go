package grpc

import (
	pb "auth/pkg/api/auth/v1"
	"context"
)

func (c *Controller) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	o := dtoRegisterFromRegisterRequest(request)

	uid, err := c.AuthUsecases.Register(o)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId: uid,
	}, nil
}
