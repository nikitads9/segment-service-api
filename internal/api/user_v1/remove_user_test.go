package user_v1

import (
	"context"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_RemoveUser(t *testing.T) {
	var (
		ctx      = context.Background()
		mock     = gomock.NewController(t)
		userId   = gofakeit.Int64()
		userErr  = errors.New(gofakeit.Phrase())
		validReq = &desc.RemoveUsertRequest{
			Id: userId,
		}
	)

	userRepoMock := userRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		userRepoMock.EXPECT().RemoveUser(ctx, userId).Return(nil).Times(1),
		userRepoMock.EXPECT().RemoveUser(ctx, userId).Return(userErr).Times(1),
	)

	api := NewMockImplementation(Implementation{
		userService: user.NewMockUserService(userRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.RemoveUser(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.RemoveUser(ctx, validReq)
		require.Error(t, err)
		require.Equal(t, err, userErr)
	})
}
