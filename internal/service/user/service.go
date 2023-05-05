package note

import (
	"context"

	"github.com/AndrewEminov/auth/internal/repository/user"

	"github.com/AndrewEminov/auth/internal/model"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, info *model.User) error
}

type service struct {
	userRepository user.Repository
}

func NewService(userRepository user.Repository) *service {
	return &service{
		userRepository: userRepository,
	}
}
