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
		Userc:     ToUserCFromService(&user.UC),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserCFromService(user *model.UserC) *desc.UserC {
	return &desc.UserC{
		Useru:    ToUserUFromService(&user.UU),
		Password: "",
	}
}

func ToUserUFromService(user *model.UserU) *desc.UserU {
	return &desc.UserU{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  desc.Role(user.Role),
	}
}

func ToUserFromDesc(user *desc.User) *model.User {
	var UpdatedAt sql.NullTime
	UpdatedAt.Valid = true
	UpdatedAt.Time = user.UpdatedAt.AsTime()
	UC := ToUserFromDescCreate(user.GetUserc())

	return &model.User{
		UC:        *UC,
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: UpdatedAt,
	}
}

func ToUserFromDescCreate(user *desc.UserC) *model.UserC {
	UU := ToUserFromDescUpdate(user.GetUseru())
	return &model.UserC{
		UU:       *UU,
		Password: user.Password,
	}
}

func ToUserFromDescUpdate(user *desc.UserU) *model.UserU {
	return &model.UserU{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  int(user.GetRole()),
	}
}
