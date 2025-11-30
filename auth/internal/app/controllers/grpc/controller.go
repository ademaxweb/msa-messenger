package grpc

import (
	"auth/internal/app/usecases"
	pb "auth/pkg/api/auth/v1"
	"errors"
)

type Deps struct {
	AuthUsecases usecases.Interface
}

type Controller struct {
	pb.UnimplementedAuthServiceServer
	Deps
}

func (d *Deps) Validate() error {
	if d.AuthUsecases == nil {
		return errors.New("deps.useCases is nil")
	}

	return nil
}

func New(d Deps) (*Controller, error) {
	if err := d.Validate(); err != nil {
		return nil, err
	}

	return &Controller{
		Deps: d,
	}, nil
}
