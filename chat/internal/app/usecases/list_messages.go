package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) ListMessages(o dto.ListMessages) ([]models.Message, error) {
	messages, err := cs.repo.GetChatMessagesByChatID(o.ID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
