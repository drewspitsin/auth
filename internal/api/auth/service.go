package auth

import (
	"github.com/drewspitsin/auth/internal/service"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
