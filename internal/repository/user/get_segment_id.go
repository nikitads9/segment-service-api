package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
