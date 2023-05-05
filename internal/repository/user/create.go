package user

import (
	"context"
	"github.com/AndrewEminov/auth/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repository) Create(ctx context.Context, user *model.User) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("username", "email", "role_id", "password", "password_confirm", "created_at", "updated_at").
		Values(
			user.Username,
			user.Email,
			user.Password,
			user.RoleId,
			user.PasswordConfirm,
			user.UpdatedAt,
			user.CreatedAt,
		).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	rows, err := r.pool.Query(ctx, query, v...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
