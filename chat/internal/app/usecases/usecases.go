package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

type (
	ChatRepository interface {
		// CreateChat Создать чат
		CreateChat(chat *models.Chat, members []uint32) (uint32, error)
		// CreateMessage создать сообщение
		CreateMessage(message *models.Message) (uint32, error)
		// GetChatByID получить чат по ID
		GetChatByID(id uint32) (*models.Chat, error)
		// GetChatMembersByChatID Получить список участников по ID чата
		GetChatMembersByChatID(id uint32) ([]uint32, error)
		// GetChatMessagesByChatID Получить список сообщений в чате по его ID
		GetChatMessagesByChatID(id uint32) ([]models.Message, error)
		// GetUserChatsByUserID Получить список чатов пользователя по его ID
		GetUserChatsByUserID(id uint32) ([]models.Chat, error)
	}
)

type Interface interface {
	// CreateDirectChat Создать 1-to-1 чат
	CreateDirectChat(o dto.CreateDirectChat) (*models.Chat, error)
	// GetChat Получить информацию о чате
	GetChat(o dto.GetChat) (*models.Chat, error)
	// ListChatMembers Получить список участников чата
	ListChatMembers(o dto.ListChatMembers) ([]uint32, error)
	// ListMessages Получить список сообщений в чате
	ListMessages(o dto.ListMessages) ([]models.Message, error)
	// ListUserChats Получить список чатов пользователя
	ListUserChats(o dto.ListUserChats) ([]models.Chat, error)
	// SendMessage Отправить сообщение
	SendMessage(o dto.SendMessage) (*models.Message, error)
	// StreamMessages Получение сообщений в реальном времени
	StreamMessages(o dto.StreamMessages) error
}

type ChatService struct {
	repo ChatRepository
}

func NewChatService(repository ChatRepository) *ChatService {
	return &ChatService{
		repo: repository,
	}
}
