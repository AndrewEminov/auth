package note

import (
	"context"
	"github.com/AndrewEminov/auth/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) Update(ctx context.Context, info *model.Info) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title", "content", "created_at").
		Values(info.Title, info.Content, info.CreatedAt).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := r.pool.Query(ctx, query, v...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	rows.Next()
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
