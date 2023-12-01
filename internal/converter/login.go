package converter

import (
	"github.com/drewspitsin/auth/internal/model"
	desc "github.com/drewspitsin/auth/pkg/auth_v1"
)

func ToUserClaimsFromLogin(req *desc.Login) *model.UserClaims {
	return &model.UserClaims{
		Username: req.GetUsername(),
	}
}
