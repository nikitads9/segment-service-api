package segment_v1

import (
	"context"
	"errors"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	segmentRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/segment_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/segment"
	desc "github.com/nikitads9/segment-service-api/pkg/segment_api"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_RemoveSegment(t *testing.T) {
	var (
		ctx       = context.Background()
		mock      = gomock.NewController(t)
		segmentId = gofakeit.Int64()

		segmentErr = errors.New(gofakeit.Phrase())

		validRequest = &desc.RemoveSegmentRequest{
			Id: segmentId,
		}
	)
	segmentRepoMock := segmentRepoMocks.NewMockRepository(mock)
	gomock.InOrder(
		segmentRepoMock.EXPECT().RemoveSegment(ctx, segmentId).Return(nil).Times(1),
		segmentRepoMock.EXPECT().RemoveSegment(ctx, segmentId).Return(segmentErr).Times(1),
	)

	api := newMockImplementation(Implementation{
		segmentService: segment.NewMockSegmentService(segmentRepoMock),
	})

	t.Run("success case", func(t *testing.T) {
		res, err := api.RemoveSegment(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.RemoveSegment(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, err, segmentErr)
	})
}
