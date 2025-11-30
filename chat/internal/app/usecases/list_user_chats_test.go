package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_ListUserChats(t *testing.T) {
	type args struct {
		o dto.ListUserChats
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

	tests := []struct {
		name        string
		args        args
		mock        func(t *testing.T) deps
		want        []models.Chat
		assertErr   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Test 1. Positive",
			args: args{
				o: dto.ListUserChats{
					UserID: 30,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetUserChatsByUserID(uint32(30)).
					Return(
						[]models.Chat{
							{
								ID:   500,
								Name: "Бесцветный чат №0",
							},
							{
								ID:   501,
								Name: "Бесцветный чат №1",
							},
						},
						nil,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want: []models.Chat{
				{
					ID:   500,
					Name: "Бесцветный чат №0",
				},
				{
					ID:   501,
					Name: "Бесцветный чат №1",
				},
			},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Positive. User has no chats",
			args: args{
				o: dto.ListUserChats{
					UserID: 2215,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetUserChatsByUserID(uint32(2215)).
					Return(
						[]models.Chat{},
						nil,
					).
					Once()

				return deps{
					chatRepo: cr,
				}
			},
			want:        []models.Chat{},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			props := tt.mock(t)

			uc := usecases.NewChatService(props.chatRepo, props.chatNameGen)

			got, err := uc.ListUserChats(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
