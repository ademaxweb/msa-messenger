package usecases_test

import (
	"chat/internal/app/models"
	"chat/internal/app/usecases"
	"chat/internal/app/usecases/dto"
	"chat/internal/app/usecases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatService_ListChatMembers(t *testing.T) {
	type args struct {
		o dto.ListChatMembers
	}

	type deps struct {
		chatRepo    usecases.ChatRepository
		chatNameGen usecases.NameGenerator
	}

	tests := []struct {
		name        string
		args        args
		mock        func(t *testing.T) deps
		want        []uint32
		assertErr   assert.ErrorAssertionFunc
		expectedErr error
	}{
		{
			name: "Test 1. Positive",
			args: args{
				o: dto.ListChatMembers{
					Id: 25,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatMembersByChatID(uint32(25)).
					Return([]uint32{1, 2, 3}, nil)

				return deps{
					chatRepo: cr,
				}
			},
			want:        []uint32{1, 2, 3},
			assertErr:   assert.NoError,
			expectedErr: nil,
		},
		{
			name: "Test 2. Negative - chat not found",
			args: args{
				o: dto.ListChatMembers{
					Id: 15,
				},
			},
			mock: func(t *testing.T) deps {
				cr := mocks.NewMockChatRepository(t)

				cr.EXPECT().
					GetChatMembersByChatID(uint32(15)).
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

			got, err := uc.ListChatMembers(tt.args.o)

			tt.assertErr(t, err)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
