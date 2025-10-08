package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) GetChat(o dto.GetChat) (*models.Chat, error) {
	c, err := cs.repo.GetChatByID(o.Id)
	if err != nil {
		return nil, err
	}

	return c, nil
}
