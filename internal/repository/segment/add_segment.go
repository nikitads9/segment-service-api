package segment

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
		return 0, ErrFailed
	}

	return id, nil
}
