package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/logger"
	"github.com/drewspitsin/auth/internal/model"
	"go.uber.org/zap"
)

func (s *serv) Create(ctx context.Context, info *model.UserCreate) (int64, error) {
	logger.Info("Creating...", zap.Int64("id", info.UserUpdate.ID))
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.authRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.authRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
