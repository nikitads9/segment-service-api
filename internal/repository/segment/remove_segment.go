package segment

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
		return ErrSegNotFound
	}

	return nil
}
