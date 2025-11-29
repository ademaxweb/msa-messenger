package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

func (us *UsersService) CreateProfile(o dto.CreateProfile) (*models.User, error) {
	u := &models.User{
		Name: o.Name,
		Bio:  o.Bio,
	}

	id, err := us.repo.Save(u)
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}
