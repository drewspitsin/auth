package converter

import (
	"github.com/drewspitsin/auth/internal/model"
	modelRepo "github.com/drewspitsin/auth/internal/repository/auth/model"
)

func ToUserFromRepo(user modelRepo.User) *model.User {
	return &model.User{
		UserCreate: ToUserCFromRepo(user.UserCreate),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

func ToUserCFromRepo(user modelRepo.UserCreate) model.UserCreate {
	return model.UserCreate{
		UserUpdate: ToUserUFromRepo(user.UserUpdate),
		Password:   user.Password,
	}
}

func ToUserUFromRepo(user modelRepo.UserUpdate) model.UserUpdate {
	return model.UserUpdate{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
