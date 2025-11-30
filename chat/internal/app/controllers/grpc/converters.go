package grpc

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases/dto"
	pb "chat/pkg/api/chat/v1"
)

func dtoGetChatFromGetChatRequest(request *pb.GetChatRequest) dto.GetChat {
	return dto.GetChat{
		ID: request.GetChatId(),
	}
}

func dtoListChatMembersFromListChatMembersRequest(request *pb.ListChatMembersRequest) dto.ListChatMembers {
	return dto.ListChatMembers{
		ID: request.GetChatId(),
	}
}

func dtoListMessagesFromListMessagesRequest(request *pb.ListMessagesRequest) dto.ListMessages {
	return dto.ListMessages{
		ID: request.GetChatId(),
	}
}

func dtoListUserChatsFromListUserChatsRequest(request *pb.ListUserChatsRequest) dto.ListUserChats {
	return dto.ListUserChats{
		UserID: request.GetUserId(),
	}
}

func pbChatFromChatModel(chat *models.Chat) *pb.Chat {
	return &pb.Chat{
		Id:          chat.ID,
		Name:        chat.Name,
		Description: chat.Description,
		AvatarUrl:   chat.AvatarURL,
	}
}

func pbMessageFromMessageModel(message *models.Message) *pb.Message {
	return &pb.Message{
		Id:     message.Id,
		ChatId: message.ChatID,
		UserId: message.UserID,
		Text:   message.Text,
	}
}
