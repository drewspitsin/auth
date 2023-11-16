package login

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/drewspitsin/auth/internal/model"
	"github.com/drewspitsin/auth/internal/utils"
)

func (s *serverAuth) GetAccessToken(ctx context.Context, token string) (string, error) {
	claims, err := utils.VerifyToken(token, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	accessToken, err := utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
