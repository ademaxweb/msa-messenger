package usecases

import "chat/internal/app/usecases/dto"

func (cs *ChatService) ListChatMembers(o dto.ListChatMembers) ([]uint32, error) {
	members, err := cs.repo.GetChatMembersByChatID(o.Id)
	if err != nil {
		return nil, err
	}

	return members, nil
}
