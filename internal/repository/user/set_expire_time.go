package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/model"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

func (r *repository) SetExpireTime(ctx context.Context, mod *model.SetExpireTimeInfo) error {
	segmentId, err := r.GetSegmentId(ctx, mod.Slug)
	if err != nil {
		return err
	}
	if segmentId == 0 {
		return ErrNotFound
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
