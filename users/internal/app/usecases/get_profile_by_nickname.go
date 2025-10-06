package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

func (us *UsersService) GetProfileByNickname(o dto.GetProfileByNickname) (*models.User, error) {
	u, err := us.repo.GetByNickname(o.Name)
	if err != nil {
		return nil, err
	}

	return u, nil
}
