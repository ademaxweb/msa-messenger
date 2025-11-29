package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_GetChat(t *testing.T) {
	type args struct {
		o dto.GetChat
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

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
				o: dto.GetChat{
					ID: 15,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatByID(uint32(15)).
					Return(
						&models.Chat{
							ID:          15,
							Name:        "Бесцветный чат №0",
							Description: "just a test chat",
							AvatarURL:   "123.png",
						},
						nil,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want: &models.Chat{
				ID:          15,
				Name:        "Бесцветный чат №0",
				Description: "just a test chat",
				AvatarURL:   "123.png",
			},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Negative - chat not found",
			args: args{
				o: dto.GetChat{
					ID: 15,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatByID(uint32(15)).
					Return(
						nil,
						models.ChatNotFound,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want:        nil,
			assertErr:   assert.Error,
			expectedErr: models.ChatNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			props := tt.mock(t)

			uc := usecases.NewChatService(props.chatRepo, props.chatNameGen)

			got, err := uc.GetChat(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
