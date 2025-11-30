package usecases

import (
	"auth/internal/app/models"
	"auth/internal/app/usecases/dto"
)

func (s *AuthService) Refresh(o dto.Refresh) (*models.Token, error) {
	return &models.Token{}, nil
}
