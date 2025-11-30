package usecases

import (
	"users/internal/app/models"
	"users/internal/app/usecases/dto"
)

type (
	UserRepository interface {
		Save(o *models.User) (uint32, error)
		GetByID(uint32) (*models.User, error)
		GetByNickname(string) (*models.User, error)
		SearchByNickname(string) ([]models.User, error)
		Update(*models.User) (*models.User, error)
	}
)

type Interface interface {
	// CreateProfile создание профиля
	CreateProfile(o dto.CreateProfile) (*models.User, error)

	// GetProfileByID получение профиля по ID
	GetProfileByID(o dto.GetProfileByID) (*models.User, error)

	// GetProfileByNickname получения профиля по ник-нейму
	GetProfileByNickname(o dto.GetProfileByNickname) (*models.User, error)

	// SearchProfileByNickname поиск профиля по ник-нейму
	SearchProfileByNickname(o dto.SearchProfileByNickname) ([]models.User, error)

	// UpdateProfile обновление информации в профиле
	UpdateProfile(o dto.UpdateProfile) (*models.User, error)
}

type UsersService struct {
	repo UserRepository
}

func NewUsersService(repo UserRepository) *UsersService {
	return &UsersService{
		repo: repo,
	}
}
