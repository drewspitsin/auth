package auth

import (
	"context"

	"github.com/drewspitsin/auth/internal/converter"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := i.authService.Delete(ctx, converter.ToUserFromDescDelete(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
