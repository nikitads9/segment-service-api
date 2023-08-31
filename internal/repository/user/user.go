package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no segment with this id")

type Repository interface {
	AddToSegment(ctx context.Context, slug string, id int64) error
	RemoveFromSegment(ctx context.Context, slug string, id int64) error
	GetSegments(ctx context.Context, id int64) ([]string, error)
	AddUser(ctx context.Context, userName string) (int64, error)
	GetUser(ctx context.Context, id int64) (string, error)
	RemoveUser(ctx context.Context, id int64) error
}
type repository struct {
	client db.Client
}

func NewUserRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

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

func (r *repository) AddToSegment(ctx context.Context, slug string, id int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return errNotFound
	}

	builder := sq.Insert(table.JunctionTable).
		Columns("user_id", "segment_id").
		Values(id, segmentId).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.AddToSegment",
		QueryRaw: query + " on conflict (user_id, segment_id) do nothing",
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) RemoveFromSegment(ctx context.Context, slug string, id int64) error {
	segmentId, err := r.GetSegmentId(ctx, slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return errNotFound
	}

	builder := sq.Delete(table.JunctionTable).
		Where(sq.And{
			sq.Eq{"user_id": id},
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
		return err
	}

	return nil
}

func (r *repository) GetSegments(ctx context.Context, id int64) ([]string, error) {
	builder := sq.Select("slug").
		From(table.JunctionTable).
		Join(table.SegmentTable + " ON segment_id=id").
		Where(sq.Eq{"user_id": id}).
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
		return "", err
	}

	return userName, nil
}

func (r *repository) AddUser(ctx context.Context, userName string) (int64, error) {
	builder := sq.Insert(table.UserTable).
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
		return 0, err
	}

	return id, nil
}

func (r *repository) RemoveUser(ctx context.Context, id int64) error {
	builder := sq.Delete(table.UserTable).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.RemoveUser",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil

}
