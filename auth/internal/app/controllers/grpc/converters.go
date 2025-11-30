package grpc

import (
	"auth/internal/app/usecases/dto"
	pb "auth/pkg/api/auth/v1"
)

func dtoLoginFromLoginRequest(req *pb.LoginRequest) dto.Login {
	return dto.Login{
		Credentials: dto.Credentials{
			Email:    req.GetCredentials().GetEmail(),
			Password: req.GetCredentials().GetPassword(),
		},
	}
}

func dtoRegisterFromRegisterRequest(req *pb.RegisterRequest) dto.Register {
	return dto.Register{
		Credentials: dto.Credentials{
			Email:    req.GetCredentials().GetEmail(),
			Password: req.GetCredentials().GetPassword(),
		},
	}
}

func dtoRefreshFromRefreshRequest(req *pb.RefreshRequest) dto.Refresh {
	return dto.Refresh{
		Token: req.GetRefreshToken(),
	}
}
