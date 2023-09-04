package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
		return 0, ErrFailed

	}

	return id, nil
}
