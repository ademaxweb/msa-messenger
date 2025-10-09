package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_ListMessages(t *testing.T) {
	type args struct {
		o dto.ListMessages
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

	tests := []struct {
		name        string
		args        args
		mock        func(t *testing.T) deps
		want        []models.Message
		assertErr   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Test 1. Positive",
			args: args{
				o: dto.ListMessages{
					Id: 15,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatMessagesByChatID(uint32(15)).
					Return(
						[]models.Message{
							{
								Id:     1,
								ChatID: 500,
								UserID: 30,
								Text:   "Привет!",
							},
							{
								Id:     2,
								ChatID: 500,
								UserID: 40,
								Text:   "Hi!",
							},
						},
						nil,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want: []models.Message{
				{
					Id:     1,
					ChatID: 500,
					UserID: 30,
					Text:   "Привет!",
				},
				{
					Id:     2,
					ChatID: 500,
					UserID: 40,
					Text:   "Hi!",
				},
			},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Negative - chat not found",
			args: args{
				o: dto.ListMessages{
					Id: 15,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatMessagesByChatID(uint32(15)).
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

			got, err := uc.ListMessages(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
