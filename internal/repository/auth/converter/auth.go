package converter

import (
	"github.com/drewspitsin/auth/internal/model"
	modelRepo "github.com/drewspitsin/auth/internal/repository/auth/model"
)

func ToUserFromRepo(user modelRepo.User) *model.User {
	return &model.User{
		UC:        ToUserCFromRepo(user.UC),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserCFromRepo(user modelRepo.UserC) model.UserC {
	return model.UserC{
		UU:       ToUserUFromRepo(user.UU),
		Password: user.Password,
	}
}

func ToUserUFromRepo(user modelRepo.UserU) model.UserU {
	return model.UserU{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
