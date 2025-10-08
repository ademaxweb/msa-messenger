package usecases

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
	"fmt"
	"math/rand"
)

func generateChatName() string {
	colors := []string{"Красный", "Зелёный", "Жёлтый", "Розовый", "Голубой", "Оранжевый"}

	return fmt.Sprintf("%s чат №%d", colors[rand.Intn(len(colors))], rand.Intn(10000)+1)
}

func (cs *ChatService) CreateDirectChat(o dto.CreateDirectChat) (*models.Chat, error) {
	c := &models.Chat{
		Name:        generateChatName(),
		Description: "", // TODO add chat description logic
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
