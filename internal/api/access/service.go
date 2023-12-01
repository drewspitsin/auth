package access

import (
	"github.com/drewspitsin/auth/internal/service"
	desc "github.com/drewspitsin/auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
