package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

func (us *UsersService) UpdateProfile(o dto.UpdateProfile) (*models.User, error) {
	u := &models.User{
		ID:        o.Id,
		Name:      o.Name,
		Bio:       o.Bio,
		AvatarUrl: o.AvatarUrl,
	}

	newU, err := us.repo.Update(u)
	if err != nil {
		return nil, err
	}

	return newU, nil
}
