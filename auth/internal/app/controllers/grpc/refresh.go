package grpc

import (
	pb "auth/pkg/api/auth/v1"
	"context"
)

func (c *Controller) Refresh(ctx context.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	o := dtoRefreshFromRefreshRequest(request)

	t, err := c.AuthUsecases.Refresh(o)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshResponse{
		Token: &pb.Token{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			UserId:       t.UserID,
		},
	}, nil
}
