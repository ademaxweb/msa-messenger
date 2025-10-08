package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) ListUserChats(o dto.ListUserChats) ([]models.Chat, error) {
	chats, err := cs.repo.GetUserChatsByUserID(o.UserId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}
