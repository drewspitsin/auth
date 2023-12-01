package login

import (
	"github.com/drewspitsin/auth/internal/service"
	desc "github.com/drewspitsin/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	loginService service.LoginService
}

func NewImplementation(loginService service.LoginService) *Implementation {
	return &Implementation{
		loginService: loginService,
	}
}
