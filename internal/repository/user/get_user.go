package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nikitads9/segment-service-api/internal/client/db"
	"github.com/nikitads9/segment-service-api/internal/repository/table"
)

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
		return "", ErrFailed
	}

	return userName, nil
}
