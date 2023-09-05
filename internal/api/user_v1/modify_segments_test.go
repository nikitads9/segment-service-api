package user_v1

import (
	"context"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	pgxV5 "github.com/jackc/pgx/v5"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	dbMocks "github.com/nikitads9/segment-service-api/internal/client/db/mocks_db"
	txMocks "github.com/nikitads9/segment-service-api/internal/client/db/mocks_tx"
	"github.com/nikitads9/segment-service-api/internal/client/db/transaction"
	userRepoMocks "github.com/nikitads9/segment-service-api/internal/repository/mocks/user_mocks"
	"github.com/nikitads9/segment-service-api/internal/service/user"
	desc "github.com/nikitads9/segment-service-api/pkg/user_api"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Test_ModifySegments(t *testing.T) {
	var (
		ctx  = context.Background()
		mock = gomock.NewController(t)

		userId    = gofakeit.Int64()
		slug1     = gofakeit.BeerStyle()
		slug2     = gofakeit.BeerStyle()
		userErr   = errors.New(gofakeit.Phrase())
		txErr     = errors.Wrap(userErr, "failed executing code inside transaction")
		txErrText = txErr.Error()
	)

	userRepoMock := userRepoMocks.NewMockRepository(mock)
	dbMock := dbMocks.NewMockDB(mock)
	txMock := txMocks.NewMockTx(mock)
	txManagerMock := transaction.NewMockTransactionManager(dbMock)
	txCtx := db.GetContextTx(ctx, txMock)
	api := NewMockImplementation(Implementation{
		userService: user.NewMockUserService(userRepoMock, txManagerMock)})

	t.Run("add one segment remove one success case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug1, userId).Return(nil).Times(1)
		userRepoMock.EXPECT().RemoveFromSegment(txCtx, slug1, userId).Return(nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1},
			SlugsToRemove: []string{slug1},
		})
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("add two segments remove one of them success case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug1, userId).Return(nil).Times(1)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug2, userId).Return(nil).Times(1)
		userRepoMock.EXPECT().RemoveFromSegment(txCtx, slug2, userId).Return(nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1, slug2},
			SlugsToRemove: []string{slug2},
		})
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("add no segments remove one segment success case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().RemoveFromSegment(txCtx, slug1, userId).Return(nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    nil,
			SlugsToRemove: []string{slug1},
		})
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("add one segment remove no segments success case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug1, userId).Return(nil).Times(1)
		txMock.EXPECT().Commit(txCtx).Return(nil)

		res, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1},
			SlugsToRemove: nil,
		})
		require.Nil(t, err)
		require.Equal(t, res, &emptypb.Empty{})
	})

	t.Run("double insert error case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug1, userId).Return(nil).Times(1)
		userRepoMock.EXPECT().AddToSegment(txCtx, slug1, userId).Return(userErr).Times(1)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    []string{slug1, slug1},
			SlugsToRemove: nil,
		})
		require.Error(t, err)
		require.Equal(t, txErrText, err.Error())
	})

	t.Run("double remove error case", func(t *testing.T) {
		dbMock.EXPECT().BeginTx(ctx, pgxV5.TxOptions{IsoLevel: pgxV5.ReadCommitted}).Return(txMock, nil)
		userRepoMock.EXPECT().RemoveFromSegment(txCtx, slug1, userId).Return(nil).Times(1)
		userRepoMock.EXPECT().RemoveFromSegment(txCtx, slug1, userId).Return(userErr).Times(1)
		txMock.EXPECT().Rollback(txCtx).Return(nil)

		_, err := api.ModifySegments(ctx, &desc.ModifySegmentsRequest{
			Id:            userId,
			SlugsToAdd:    nil,
			SlugsToRemove: []string{slug1, slug1},
		})
		require.Error(t, err)
		require.Equal(t, txErrText, err.Error())
	})

}
