package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/logger"
	"go.uber.org/zap"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	logger.Info("Deleating...", zap.Int64("id", id))
	err := s.authRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
