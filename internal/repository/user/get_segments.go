package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
