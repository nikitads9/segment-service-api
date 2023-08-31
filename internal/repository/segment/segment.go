package segment

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/segment_service_repository.go -package=mocks . Repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "there is no segment with this id")
var errFailed = status.Error(codes.InvalidArgument, "the operation failed")

type Repository interface {
	AddSegment(ctx context.Context, slug string) (int64, error)
	RemoveSegment(ctx context.Context, id int64) error
}
type repository struct {
	client db.Client
}

func NewSegmentRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) AddSegment(ctx context.Context, slug string) (int64, error) {
	builder := sq.Insert(table.SegmentTable).
		PlaceholderFormat(sq.Dollar).
		Columns("slug").
		Values(slug).
		Suffix("ON CONFLICT (slug) DO NOTHING returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "segment_repository.AddSegment",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, errFailed
	}

	return id, nil
}

func (r *repository) RemoveSegment(ctx context.Context, id int64) error {
	builder := sq.Delete(table.SegmentTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "segment_repository.RemoveSegment",
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
