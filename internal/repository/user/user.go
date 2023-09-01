package user

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/user_mocks/user_service_repository.go -package=user_mocks . Repository

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/model"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no instance with this name or id")
var errFailed = status.Error(codes.InvalidArgument, "the operation failed")

type Repository interface {
	AddToSegment(ctx context.Context, slug string, userId int64) error
	RemoveFromSegment(ctx context.Context, slug string, userId int64) error
	GetSegments(ctx context.Context, userId int64) ([]string, error)
	SetExpireTime(ctx context.Context, mod *model.SetExpireTimeInfo) error
	AddUser(ctx context.Context, userName string) (int64, error)
	GetUser(ctx context.Context, userId int64) (string, error)
	RemoveUser(ctx context.Context, userId int64) error
	GetUserHistoryCsv(ctx context.Context, userId int64) (bytes.Buffer, error)
}

type repository struct {
	client db.Client
}

func NewUserRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

// Метод для совершения вложенного запроса поиска id сегмента по его названию
func (r *repository) GetSegmentId(ctx context.Context, slug string) (int64, error) {
	var segmentId int64

	subQ := sq.Select("id").
		From(table.SegmentTable).
		Where(sq.Eq{"slug": slug}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := subQ.ToSql()
	if err != nil {
		return segmentId, err
	}

	q := db.Query{
		Name:     "user_repository.GetSegmentId",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return segmentId, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(&segmentId)
		if err != nil {
			return segmentId, err
		}
	}
	return segmentId, nil
}

// Метод для добавления пользователя (id) в указанный сегмент(slug)
func (r *repository) AddToSegment(ctx context.Context, slug string, userId int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return errNotFound
	}

	builder := sq.Insert(table.JunctionTable).
		Columns("junction_id", "user_id", "segment_id", "state", "added_at").
		Values(uuid.New(), userId, segmentId, true, time.Now().UTC()).
		PlaceholderFormat(sq.Dollar).
		Suffix("ON CONFLICT (user_id, segment_id, state) DO NOTHING")

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

func (r *repository) RemoveFromSegment(ctx context.Context, slug string, userId int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return errNotFound
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

func (r *repository) GetSegments(ctx context.Context, userId int64) ([]string, error) {
	builder := sq.Select("slug").
		From(table.JunctionTable).
		Join(table.SegmentTable + " ON segment_id=id").
		Where(sq.And{
			sq.Eq{"user_id": userId},
			sq.Eq{"state": true},
		}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetSegments",
		QueryRaw: query,
	}

	var segments []string
	err = r.client.DB().SelectContext(ctx, &segments, q, args...)
	if err != nil {
		return nil, err
	}

	return segments, nil
}

func (r *repository) SetExpireTime(ctx context.Context, mod *model.SetExpireTimeInfo) error {
	segmentId, err := r.GetSegmentId(ctx, mod.Slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return errNotFound
	}

	builder := sq.Update(table.JunctionTable).
		Set("time_of_expire", mod.ExpireTime).
		Where(sq.And{
			sq.Eq{"user_id": mod.UserId},
			sq.Eq{"segment_id": segmentId},
		}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.SetExpireTime",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, id int64) (string, error) {
	builder := sq.Select("user_name").
		From(table.UserTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "user_repository.GetUser",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return "", err
	}

	var userName string
	row.Next()
	err = row.Scan(&userName)
	if err != nil {
		return "", errFailed
	}

	return userName, nil
}

func (r *repository) AddUser(ctx context.Context, userName string) (int64, error) {
	builder := sq.Insert(table.UserTable).
		Columns("username").
		Values(userName).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.AddUser",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, errFailed

	}

	return id, nil
}

func (r *repository) RemoveUser(ctx context.Context, userId int64) error {
	builder := sq.Delete(table.UserTable).
		Where(sq.Eq{"id": userId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.RemoveUser",
		QueryRaw: query,
	}

	result, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errNotFound
	}

	return nil

}

func (r *repository) GetUserHistoryCsv(ctx context.Context, userId int64) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	builder := sq.Select("slug", "added_at", "time_of_expire").
		From(table.JunctionTable).Join(table.SegmentTable + " ON segment_id=id").
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return buffer, err
	}

	q := db.Query{
		Name:     "user_repository.GetUserHistoryCsv",
		QueryRaw: query,
	}

	var dest []model.HistoryLine
	err = r.client.DB().SelectContext(ctx, &dest, q, args...)
	if err != nil {
		return buffer, err
	}

	writer := csv.NewWriter(&buffer)
	for _, elem := range dest {
		writer.Write(elem.ToStringArray())
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return buffer, err
	}

	return buffer, nil
}
