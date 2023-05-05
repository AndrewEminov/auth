package note

import (
	"context"

	"github.com/AndrewEminov/auth/internal/model"
)

func (s *service) Update(ctx context.Context, info *model.Info) (int64, error) {
	id, err := s.noteRepository.Update(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
