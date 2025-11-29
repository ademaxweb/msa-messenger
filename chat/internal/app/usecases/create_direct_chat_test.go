package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	nameGenMocks "chat/pkg/namegen/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_CreateDirectChat(t *testing.T) {
	type args struct {
		o dto.CreateDirectChat
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

	chatName := "Бесцветный чат №0"

	tests := []struct {
		name        string
		args        args
		mock        func(t *testing.T) deps
		want        *models.Chat
		assertErr   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Test 1. Positive",
			args: args{
				o: dto.CreateDirectChat{
					SenderID:    30,
					RecipientID: 555,
				},
			},
			mock: func(t *testing.T) deps {
				ng := nameGenMocks.NewMockGenerator(chatName)

				cr := mocks.NewMockChatRepository(t)

				members := []uint32{30, 555}

				cr.EXPECT().
					CreateChat(
						&models.Chat{
							Name: ng.Generate(),
						},
						members,
					).
					Return(uint32(666), nil).
					Once()

				return deps{
					chatRepo:    cr,
					chatNameGen: ng,
				}
			},
			want: &models.Chat{
				ID:   666,
				Name: chatName,
			},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Negative. Cannot create chat with yourself",
			args: args{
				o: dto.CreateDirectChat{
					SenderID:    30,
					RecipientID: 30,
				},
			},
			mock: func(t *testing.T) deps {
				ng := nameGenMocks.NewMockGenerator(chatName)
				cr := mocks.NewMockChatRepository(t)

				return deps{
					chatRepo:    cr,
					chatNameGen: ng,
				}
			},
			want:        nil,
			assertErr:   assert.Error,
			expectedErr: models.CannotChatToYourself,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			props := tt.mock(t)

			uc := usecases.NewChatService(props.chatRepo, props.chatNameGen)

			got, err := uc.CreateDirectChat(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
