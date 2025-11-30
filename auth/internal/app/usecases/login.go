package usecases

import (
	"auth/internal/app/models"
	"auth/internal/app/usecases/dto"
)

func (s *AuthService) Login(o dto.Login) (*models.Token, error) {
	return &models.Token{}, nil
}
