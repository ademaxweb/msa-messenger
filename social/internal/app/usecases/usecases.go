package usecases

import (
	"social/internal/app/models"
	"social/internal/app/usecases/dto"
)

type (
	RequestsRepository interface {
		// CreateRequest Создание заявки
		CreateRequest(request *models.FriendRequest) (uint32, error)

		// UpdateRequest Обновление заявки по её ID
		UpdateRequest(request *models.FriendRequest) (*models.FriendRequest, error)

		// UpdateRequestByUsers Обновление заявки по ID её участников (recipient_id и sender_id)
		UpdateRequestByUsers(request *models.FriendRequest) (*models.FriendRequest, error)

		// GetUserRequests Получение списка заявок пользователя
		GetUserRequests(filters *models.FriendRequest) ([]models.FriendRequest, error)
	}
)

type Interface interface {
	// UpdateRequestStatus Обновление статуса заявки по её ID
	UpdateRequestStatus(o dto.UpdateRequestStatus) (*models.FriendRequest, error)

	// ListRequests Получение списка заявок у пользователя
	ListRequests(o dto.ListRequests) ([]models.FriendRequest, error)

	// ListFriends Получение списка друзей у пользователя
	ListFriends(o dto.ListFriends) ([]uint32, error)

	// RemoveFriend Удаление пользователя из списка друзей
	RemoveFriend(o dto.RemoveFriend) error

	// SendFriendRequest Отправление заявки в друзья
	SendFriendRequest(o dto.SendFriendRequest) (*models.FriendRequest, error)
}

type RequestsService struct {
	repo RequestsRepository
}

func NewRequestsService(repo RequestsRepository) *RequestsService {
	return &RequestsService{repo: repo}
}
