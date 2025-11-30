package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

func (us *UsersService) GetProfileByID(o dto.GetProfileByID) (*models.User, error) {
	u, err := us.repo.GetByID(o.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}
