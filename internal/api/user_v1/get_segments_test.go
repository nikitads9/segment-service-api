package user_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/stretchr/testify/require"
)

func Test_GetSegments(t *testing.T) {
	var (
		ctx    = context.Background()
		mock   = gomock.NewController(t)
		userId = gofakeit.Int64()

		userErr = errors.New(gofakeit.Phrase())
		slug1   = gofakeit.BeerName()
		slug2   = gofakeit.BeerName()

		validRequest = &desc.GetSegmentsRequest{
			Id: userId,
		}

		validResponse = &desc.GetSegmentsResponse{
			Slugs: []string{slug1, slug2},
		}
	)

	userRepoMock := userRepoMocks.NewMockRepository(mock)
	api := NewMockImplementation(Implementation{
		userService: user.NewMockUserService(userRepoMock)})

	gomock.InOrder(
		userRepoMock.EXPECT().GetSegments(ctx, userId).Return([]string{slug1, slug2}, nil).Times(1),
		userRepoMock.EXPECT().GetSegments(ctx, userId).Return(nil, userErr).Times(1),
	)

	t.Run("success case", func(t *testing.T) {
		res, err := api.GetSegments(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, validResponse, res)
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.GetSegments(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, userErr, err)
	})

}
