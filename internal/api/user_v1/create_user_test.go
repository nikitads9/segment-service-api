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
)

func Test_AddUser(t *testing.T) {
	var (
		ctx      = context.Background()
		mock     = gomock.NewController(t)
		userId   = gofakeit.Int64()
		userName = gofakeit.BeerName()

		userErr = errors.New(gofakeit.Phrase())

		validRequest = &desc.AddUserRequest{
			UserName: userName,
		}

		validResponse = &desc.AddUserResponse{
			Id: userId,
		}
	)
	userRepoMock := userRepoMocks.NewMockRepository(mock)

	gomock.InOrder(
		userRepoMock.EXPECT().AddUser(ctx, userName).Return(userId, nil).Times(1),
		userRepoMock.EXPECT().AddUser(ctx, userName).Return(int64(0), userErr).Times(1),
	)

	api := newMockImplementation(Implementation{
		userService: user.NewMockUserService()})

	t.Run("success case", func(t *testing.T) {
		res, err := api.AddUser(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, res.GetId(), validResponse.GetId())
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.AddUser(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, err, userErr)
	})

}
