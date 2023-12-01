package access

import (
	"github.com/drewspitsin/auth/internal/client/db"
	"github.com/drewspitsin/auth/internal/repository"
	"github.com/drewspitsin/auth/internal/service"
)

type serverAccess struct {
	accessRepository repository.AccessRepository
	txManager        db.TxManager
}

func NewService(
	accessRepository repository.AccessRepository,
	txManager db.TxManager,
) service.AccessService {
	return &serverAccess{
		accessRepository: accessRepository,
		txManager:        txManager,
	}
}
