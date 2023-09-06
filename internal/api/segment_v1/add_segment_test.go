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
)

func Test_AddSegment(t *testing.T) {
	var (
		ctx       = context.Background()
		mock      = gomock.NewController(t)
		segmentId = gofakeit.Int64()
		slug      = gofakeit.BeerName()

		segmentErr = errors.New(gofakeit.Phrase())

		validRequest = &desc.AddSegmentRequest{
			Slug: slug,
		}

		validResponse = &desc.AddSegmentResponse{
			Id: segmentId,
		}
	)
	segmentRepoMock := segmentRepoMocks.NewMockRepository(mock)
	api := newMockImplementation(Implementation{
		segmentService: segment.NewMockSegmentService(segmentRepoMock),
	})

	gomock.InOrder(
		segmentRepoMock.EXPECT().AddSegment(ctx, slug).Return(segmentId, nil).Times(1),
		segmentRepoMock.EXPECT().AddSegment(ctx, slug).Return(int64(0), segmentErr).Times(1),
	)

	t.Run("success case", func(t *testing.T) {
		res, err := api.AddSegment(ctx, validRequest)
		require.Nil(t, err)
		require.Equal(t, validResponse.GetId(), res.GetId())
	})

	t.Run("error case", func(t *testing.T) {
		_, err := api.AddSegment(ctx, validRequest)
		require.Error(t, err)
		require.Equal(t, segmentErr, err)
	})
}
