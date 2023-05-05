package user

import (
	"context"
	"github.com/AndrewEminov/auth/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, user *model.User) error
	//Update(ctx context.Context, user *model.UserUpdate) error
	//Delete(ctx context.Context, username string) error
	//Get(ctx context.Context, username string) (model.User, error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}
