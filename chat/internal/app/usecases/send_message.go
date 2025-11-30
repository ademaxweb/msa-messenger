package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) SendMessage(o dto.SendMessage) (*models.Message, error) {
	_, err := cs.repo.GetChatByID(o.ChatID)
	if err != nil {
		return nil, err
	}

	m := &models.Message{
		ChatID: o.ChatID,
		UserID: o.SenderID,
		Text:   o.Text,
	}

	messageId, err := cs.repo.CreateMessage(m)
	if err != nil {
		return nil, err
	}

	m.Id = messageId

	return m, nil
}
