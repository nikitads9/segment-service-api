package user_v1

import (
	"context"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/nikitads9/segment-service-api/internal/model"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_SetExpireTime(t *testing.T) {
	var (
		ctx     = context.Background()
		mock    = gomock.NewController(t)
		userId  = gofakeit.Int64()
		slug    = gofakeit.CarModel()
		expTime = gofakeit.Date()
		userErr = errors.New(gofakeit.Phrase())

		validRequest = &desc.SetExpireTimeRequest{
			Id:             userId,
			Slug:           slug,
			ExpirationTime: &timestamppb.Timestamp{Seconds: expTime.Unix()},
		}

		setExpireTimeInfo = &model.SetExpireTimeInfo{
			UserId:     userId,
			Slug:       slug,
			ExpireTime: expTime,
		}
	)
	userRepoMock := userRepoMocks.NewMockRepository(mock)
	api := NewMockImplementation(Implementation{
		userService: user.NewMockUserService(userRepoMock)})

	gomock.InOrder(
		userRepoMock.EXPECT().SetExpireTime(ctx, setExpireTimeInfo).Return(nil).Times(1),
		userRepoMock.EXPECT().SetExpireTime(ctx, setExpireTimeInfo).Return(userErr).Times(1),
	)

	t.Run("success case", func(t *testing.T) {
		res, err := api.SetExpireTime(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, emptypb.Empty{}, res)
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.SetExpireTime(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, userErr, err)
	})

}
