package auth

import (
	"context"
	"log"

	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	authObj, err := i.authService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d", authObj.ID)

	return &desc.GetResponse{
		Id:        authObj.ID,
		Name:      authObj.Name,
		Email:     authObj.Email,
		Role:      desc.Role(authObj.Role),
		CreatedAt: &timestamppb.Timestamp{},
		UpdatedAt: &timestamppb.Timestamp{},
	}, nil
}
