package converter

import (
	"database/sql"

	"github.com/drewspitsin/auth/internal/model"
	desc "github.com/drewspitsin/auth/pkg/user_api_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		UserCreate: ToUserCreateFromService(&user.UserCreate),
		CreatedAt:  timestamppb.New(user.CreatedAt),
		UpdatedAt:  updatedAt,
	}
}

func ToUserCreateFromService(user *model.UserCreate) *desc.UserCreate {
	return &desc.UserCreate{
		UserUpdate: ToUserUpdateFromService(&user.UserUpdate),
		Password:   user.Password,
	}
}

func ToUserUpdateFromService(user *model.UserUpdate) *desc.UserUpdate {
	return &desc.UserUpdate{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  desc.Role(user.Role),
	}
}

func ToUserFromDesc(user *desc.User) *model.User {
	var updatedAt sql.NullTime
	if user.UpdatedAt.CheckValid() == nil {
		updatedAt.Valid = true
		updatedAt.Time = user.UpdatedAt.AsTime()
	} else {
		updatedAt.Valid = false
	}

	userCreate := ToUserFromDescCreate(user.GetUserCreate())
	return &model.User{
		UserCreate: *userCreate,
		CreatedAt:  user.CreatedAt.AsTime(),
		UpdatedAt:  updatedAt,
	}
}

func ToUserFromDescCreate(user *desc.UserCreate) *model.UserCreate {
	userUpdate := ToUserFromDescUpdate(user.GetUserUpdate())
	return &model.UserCreate{
		UserUpdate: *userUpdate,
		Password:   user.Password,
	}
}

func ToUserFromDescUpdate(user *desc.UserUpdate) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  int(user.GetRole()),
	}
}
