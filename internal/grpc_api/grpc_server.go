package grpc_api

import (
	"context"

	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserV1Server
type UserV1Server struct {
	desc.UnimplementedUserV1Server
}

// NewUserV1Server returns a new UserV1Server instance
func NewUserV1Server() *UserV1Server {
	return &UserV1Server{}
}

// Create is a method that implements the Create method of the UserV1Server
func (s *UserV1Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return nil, nil
}

// Get is a method that implements the Get method of the UserV1Server
func (s *UserV1Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	return nil, nil
}

// Update is a method that implements the Update method of the UserV1Server
func (s *UserV1Server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Delete is a method that implements the Delete method of the UserV1Server
func (s *UserV1Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
