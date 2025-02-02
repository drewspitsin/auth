package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/converter"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := i.authService.Update(ctx, converter.ToUserFromDescUpdate(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
