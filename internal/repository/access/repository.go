package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(dbClient db.Client) repository.AccessRepository {
	return &repo{db: dbClient}
}

func (r *repo) Check(ctx context.Context, endpointAddress string) error {
	return nil
}
