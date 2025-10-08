package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) CreateDirectChat(o dto.CreateDirectChat) (*models.Chat, error) {
	c := &models.Chat{
		Name:        "New chat", // TODO add chat naming logic
		Description: "",         // TODO add chat description logic
		AvatarURL:   "",
	}

	chatMembers := []uint32{o.SenderID, o.RecipientID}

	chatId, err := cs.repo.CreateChat(c, chatMembers)
	if err != nil {
		return nil, err
	}

	c.Id = chatId

	return c, nil
}
