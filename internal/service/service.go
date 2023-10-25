package service

import (
	"context"

	"github.com/drewspitsin/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.UserC) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserU) error
	Delete(ctx context.Context, id int64) error
}
