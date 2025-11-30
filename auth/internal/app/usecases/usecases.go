package usecases

import (
	"auth/internal/app/models"
	"auth/internal/app/usecases/dto"
)

type Interface interface {
	Register(o dto.Register) (uint32, error)
	Login(o dto.Login) (*models.Token, error)
	Refresh(o dto.Refresh) (*models.Token, error)
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}
