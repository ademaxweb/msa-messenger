package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

func (us *UsersService) SearchProfileByNickname(o dto.SearchProfileByNickname) ([]models.User, error) {
	users, err := us.repo.SearchByNickname(o.Query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
