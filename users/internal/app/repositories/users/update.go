package users

import "users/internal/app/models"

func (r *Repository) Update(o *models.User) (*models.User, error) {
	return o, nil
}
