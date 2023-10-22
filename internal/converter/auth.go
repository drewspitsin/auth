package converter

import (
	"database/sql"
	"time"

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
		CreatedAt: &timestamppb.Timestamp{},
		UpdatedAt: updatedAt,
	}
}

func ToUserFromDescCreate(user *desc.CreateRequest) *model.User {

	return &model.User{
		ID:        0,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      int8(user.Role),
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{},
	}
}

func ToUserFromDescUpdate(user *desc.UpdateRequest) *model.User {

	return &model.User{
		ID:        user.Id,
		Name:      user.Name.Value,
		Email:     user.Email.Value,
		Password:  "",
		Role:      int8(user.Role),
		CreatedAt: time.Now(),
		UpdatedAt: sql.NullTime{},
	}
}

func ToUserFromDescDelete(user *desc.DeleteRequest) *model.User {

	return &model.User{
		ID:        user.Id,
		Name:      "",
		Email:     "",
		Password:  "",
		Role:      0,
		CreatedAt: time.Time{},
		UpdatedAt: sql.NullTime{},
	}
}
