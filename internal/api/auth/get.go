package auth

import (
	"context"
	"log"

	"github.com/drewspitsin/auth/internal/converter"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	authObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d", authObj.UC.UU.ID)

	return &desc.GetResponse{
		Info: converter.ToUserFromService(authObj),
	}, nil
}
