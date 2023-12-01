package login

import (
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/repository"
	"github.com/drewspitsin/auth/internal/service"
)

type serverAuth struct {
	loginRepository repository.LoginRepository
	txManager       db.TxManager
}

func NewService(
	loginRepository repository.LoginRepository,
	txManager db.TxManager,
) service.LoginService {
	return &serverAuth{
		loginRepository: loginRepository,
		txManager:       txManager,
	}
}
