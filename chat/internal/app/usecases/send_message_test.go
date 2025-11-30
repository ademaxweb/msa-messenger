package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_SendMessage(t *testing.T) {
	type args struct {
		o dto.SendMessage
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

	tests := []struct {
		name        string
		args        args
		mock        func(t *testing.T) deps
		want        *models.Message
		assertErr   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Test 1. Positive",
			args: args{
				o: dto.SendMessage{
					ChatID:   500,
					SenderID: 30,
					Text:     "How are you?",
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatByID(uint32(500)).
					Return(
						&models.Chat{
							ID:   500,
							Name: "Бесцветный чат №0",
						},
						nil,
					).
					Once()

				msg := &models.Message{
					ChatID: 500,
					UserID: 30,
					Text:   "How are you?",
				}

				cr.EXPECT().
					CreateMessage(msg).
					Return(uint32(3), nil)

				return deps{
					chatRepo: cr,
				}
			},
			want: &models.Message{
				Id:     3,
				ChatID: 500,
				UserID: 30,
				Text:   "How are you?",
			},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Negative - chat not found",
			args: args{
				o: dto.SendMessage{
					ChatID:   320,
					SenderID: 40,
					Text:     "Have a nice day",
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatByID(uint32(320)).
					Return(
						nil,
						models.ErrChatNotFound,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want:        nil,
			assertErr:   assert.Error,
			expectedErr: models.ErrChatNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			props := tt.mock(t)

			uc := usecases.NewChatService(props.chatRepo, props.chatNameGen)

			got, err := uc.SendMessage(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
