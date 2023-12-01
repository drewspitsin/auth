package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/logger"
	"github.com/drewspitsin/auth/internal/model"
	"go.uber.org/zap"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	logger.Info("Creating...", zap.Int64("id", id))
	user, err := s.authRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
