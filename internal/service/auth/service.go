package auth

import (
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/repository"
	"github.com/drewspitsin/auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
	txManager      db.TxManager
}

func NewService(
	authRepository repository.AuthRepository,
	txManager db.TxManager,
) service.AuthService {
	return &serv{
		authRepository: authRepository,
		txManager:      txManager,
	}
}

func NewMockService(deps ...interface{}) service.AuthService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.AuthRepository:
			srv.authRepository = s
		}
	}

	return &srv
}
