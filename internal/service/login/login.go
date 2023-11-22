package login

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/utils"
)

const (
	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 1 * time.Minute
)

func (s *serverAuth) Login(ctx context.Context, info *model.UserClaims) (string, error) {
	r, err := s.loginRepository.GetUserRole(ctx)
	if err != nil {
		return "", nil
	}
	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: info.Username,
		Role:     r,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
