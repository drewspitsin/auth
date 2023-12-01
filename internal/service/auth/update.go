package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/logger"
	"github.com/drewspitsin/auth/internal/model"
	"go.uber.org/zap"
)

func (s *serv) Update(ctx context.Context, info *model.UserUpdate) error {
	logger.Info("Creating...", zap.Int64("id", info.ID))
	err := s.authRepository.Update(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
