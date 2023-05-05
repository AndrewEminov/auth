package note

import (
	"context"

	"github.com/AndrewEminov/auth/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.User) error {
	err := s.userRepository.Create(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
