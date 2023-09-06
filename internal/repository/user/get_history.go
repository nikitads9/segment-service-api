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

func (r *repository) GetUserHistoryCsv(ctx context.Context, userId int64) (*bytes.Buffer, error) {
	var buffer bytes.Buffer

	builder := sq.Select("slug", "added_at", "time_of_expire").
		From(table.JunctionTable).Join(table.SegmentTable + " ON segment_id=id").
		Where(sq.Eq{"user_id": userId}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetUserHistoryCsv",
		QueryRaw: query,
	}

	var dest []model.HistoryLine
	err = r.client.DB().SelectContext(ctx, &dest, q, args...)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(&buffer)
	err = writer.Write([]string{"slug", "added_at", "time_of_expire"})
	if err != nil {
		return nil, err
	}

	for _, elem := range dest {
		err = writer.Write(elem.ToStringArray())
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return &buffer, nil
}
