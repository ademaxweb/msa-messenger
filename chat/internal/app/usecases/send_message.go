package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) SendMessage(o dto.SendMessage) (*models.Message, error) {
	_, err := cs.repo.GetChatByID(o.ChatId)
	if err != nil {
		return nil, err
	}

	m := &models.Message{
		ChatID: o.ChatId,
		UserID: o.SenderId,
		Text:   o.Text,
	}

	messageId, err := cs.repo.CreateMessage(m)
	if err != nil {
		return nil, err
	}

	m.Id = messageId

	return m, nil
}
