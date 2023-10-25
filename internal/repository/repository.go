package repository

import (
	"context"

	"github.com/drewspitsin/auth/internal/model"
)

type AuthRepository interface {
	Create(ctx context.Context, info *model.UserC) (int64, error)
	Get(ctx context.Context, userTableID int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserU) error
	Delete(ctx context.Context, id int64) error
}
