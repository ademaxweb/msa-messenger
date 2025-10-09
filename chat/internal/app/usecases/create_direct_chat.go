package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
)

func (cs *ChatService) CreateDirectChat(o dto.CreateDirectChat) (*models.Chat, error) {
	if o.SenderID == o.RecipientID {
		return nil, models.CannotChatToYourself
	}

	chatMembers := []uint32{o.SenderID, o.RecipientID}
	c := &models.Chat{
		Name:        cs.nameGen.Generate(),
		Description: "", // TODO add chat description logic
		AvatarURL:   "",
	}

	chatId, err := cs.repo.CreateChat(c, chatMembers)
	if err != nil {
		return nil, err
	}

	c.Id = chatId

	return c, nil
}
