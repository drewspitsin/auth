package login

import (
	"context"
	"fmt"

	desc "github.com/drewspitsin/auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {

	t := req.GetRefreshToken()
	fmt.Println("TEST1: ", t)
	Obj, err := i.loginService.GetRefreshToken(ctx, t)
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: Obj,
	}, nil
}
