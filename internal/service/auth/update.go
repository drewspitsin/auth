package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.UserUpdate) error {
	err := s.authRepository.Update(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
