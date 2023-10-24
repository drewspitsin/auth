package service

import (
	"context"

	"github.com/drewspitsin/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, info *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
