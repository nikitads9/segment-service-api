package user

import (
	"bytes"
	"context"
	"encoding/csv"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/model"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
