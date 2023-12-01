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

type AccessService interface {
	Check(ctx context.Context, endpointAddress string) error
}

type LoginService interface {
	Login(ctx context.Context, info *model.UserClaims) (string, error)
	GetAccessToken(ctx context.Context, token string) (string, error)
	GetRefreshToken(ctx context.Context, token string) (string, error)
}
