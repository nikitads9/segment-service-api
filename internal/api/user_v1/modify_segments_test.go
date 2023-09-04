package user_v1

import (
	"context"
	"errors"
	_ "errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_ModifySegments(t *testing.T) {
	var (
		ctx      = context.Background()
		mock     = gomock.NewController(t)
		userId   = gofakeit.Int64()
		slug1    = gofakeit.BeerStyle()
		slug2    = gofakeit.BeerStyle()
		slug3    = gofakeit.BeerStyle()
		slug4    = gofakeit.BeerStyle()
		userErr  = errors.New(gofakeit.Phrase())
		validReq = &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1, slug2},
			SlugsToRemove: []string{slug2},
		}
		invalidReq = &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1},
			SlugsToRemove: []string{slug2},
		}
		noAddReq = &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    nil,
			SlugsToRemove: []string{slug3, slug4},
		}
		noRemReq = &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1, slug2},
			SlugsToRemove: nil,
		}
	)

	userRepoMock := userRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		userRepoMock.EXPECT().AddToSegment(ctx, slug1, userId).Return(nil).Times(1),
		userRepoMock.EXPECT().AddToSegment(ctx, slug2, userId).Return(nil).Times(1),
		userRepoMock.EXPECT().RemoveFromSegment(ctx, slug2, userId).Return(nil).Times(1),
		userRepoMock.EXPECT().AddToSegment(ctx, slug1, userId).Return(userErr).Times(1),
		userRepoMock.EXPECT().RemoveFromSegment(ctx, slug2, userId).Return(userErr).Times(1),
	)

	api := newMockImplementation(Implementation{
		userService: user.NewMockUserService()})

	t.Run("success case", func(t *testing.T) {
		res, err := api.ModifySegments(ctx, validReq)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.ModifySegments(ctx, invalidReq)
		require.Error(t, err)
		require.Equal(t, err, userErr)
	})

	t.Run("no slice to add case", func(t *testing.T) {
		res, err := api.ModifySegments(ctx, noAddReq)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("no slice to remove case", func(t *testing.T) {
		res, err := api.ModifySegments(ctx, noRemReq)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})
}
