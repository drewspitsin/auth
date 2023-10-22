package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/model"
)

func (s *serv) Delete(ctx context.Context, info *model.User) error {
	err := s.authRepository.Delete(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
