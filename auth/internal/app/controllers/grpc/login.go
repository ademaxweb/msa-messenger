package grpc

import (
	pb "auth/pkg/api/auth/v1"
	"context"
)

func (c *Controller) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	o := dtoLoginFromLoginRequest(request)
	
	t, err := c.AuthUsecases.Login(o)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: &pb.Token{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			UserId:       t.UserID,
		},
	}, nil
}
