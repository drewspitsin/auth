package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/drewspitsin/auth/internal/model"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
)

func ToUserFromService(user *model.User) *desc.User {

	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      desc.Role(user.Role),
		UpdatedAt: updatedAt,
	}
}

func ToUserFromDescCreate(user *desc.CreateRequest) *model.UserCreate {
	return &model.UserCreate{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     int(user.Role),
	}
}

func ToUserFromDescUpdate(user *desc.UpdateRequest) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    user.Id,
		Name:  user.Name.Value,
		Email: user.Email.Value,
		Role:  int(user.Role),
	}
}

func ToUserFromDescDelete(user *desc.DeleteRequest) int64 {
	return user.Id
}
