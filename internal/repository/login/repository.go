package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.LoginRepository {
	return &repo{db: dbClient}
}

func (r *repo) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	return "", nil
}
func (r *repo) GetAccessToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
func (r *repo) GetRefreshToken(ctx context.Context, token string) (string, error) {
	return "", nil
}
func (r *repo) GetUserRole(ctx context.Context) (string, error) {
	return "admin", nil
}
