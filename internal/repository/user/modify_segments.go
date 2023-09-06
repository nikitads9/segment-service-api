package user

import (
	"context"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Метод для добавления пользователя (id) в указанный сегмент(slug)
func (r *repository) AddToSegment(ctx context.Context, slug string, userId int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return ErrNotFound
	}

	isActive, err := r.CheckForDuplicates(ctx, segmentId, userId)
	if err != nil {
		return err
	}
	if isActive == int64(1) {
		return ErrFailed
	}

	builder := sq.Insert(table.JunctionTable).
		Columns("junction_id", "user_id", "segment_id", "state", "added_at").
		Values(uuid.New(), userId, segmentId, true, time.Now().UTC()).
		PlaceholderFormat(sq.Dollar).
		Suffix("ON CONFLICT (user_id, segment_id, state, added_at) DO NOTHING")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.AddToSegment",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CheckForDuplicates(ctx context.Context, segmentId int64, userId int64) (int64, error) {
	check := sq.Select("1").
		From(table.JunctionTable).
		Where(
			sq.And{
				sq.Eq{"segment_id": segmentId},
				sq.Eq{"user_id": userId},
				sq.Eq{"state": true}}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := check.ToSql()
	if err != nil {
		return int64(0), err
	}

	q := db.Query{
		Name:     "user_repository.RemoveSubquery",
		QueryRaw: query,
	}

	var isActive int64
	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return int64(0), err
	}

	row.Next()
	_ = row.Scan(&isActive)

	return isActive, nil
}

func (r *repository) RemoveFromSegment(ctx context.Context, slug string, userId int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return ErrNotFound
	}

	isActive, err := r.CheckForDuplicates(ctx, segmentId, userId)
	if err != nil {
		return err
	}
	if isActive != int64(1) {
		return ErrNotFound
	}

	builder := sq.Update(table.JunctionTable).
		Set("state", false).
		Set("time_of_expire", time.Now().UTC()).
		Where(sq.And{
			sq.Eq{"user_id": userId},
			sq.Eq{"segment_id": segmentId},
		}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.RemoveFromSegment",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return status.Error(codes.InvalidArgument, pgErr.Message)
		}
		return err
	}

	return nil
}
