package note

import (
	"context"

	"github.com/AndrewEminov/auth/internal/model"
	"github.com/AndrewEminov/auth/internal/repository/note"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, info *model.Info) (int64, error)
	Update(ctx context.Context, info *model.Info) (int64, error)
}

type service struct {
	noteRepository note.Repository
}

func NewService(noteRepository note.Repository) *service {
	return &service{
		noteRepository: noteRepository,
	}
}
