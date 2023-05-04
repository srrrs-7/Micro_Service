package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/srrrs-7/micro_service.git/statistics/util/env"
)

func generateToken(userId string) (string, error) {
	// シークレットキー
	env, err := env.LoadEnv()
	if err != nil {
		return "", err
	}

	// クレームを定義
	claims := jwtConfig{
		UserID: userId,
		Role:   "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // 3時間後にトークンが無効になる
		},
	}

	// トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(env.ClientSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (*jwtConfig, error) {
	// シークレットキー
	env, err := env.LoadEnv()
	if err != nil {
		return nil, err
	}
	// トークンを検証
	token, err := jwt.ParseWithClaims(tokenString, &jwtConfig{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.ClientSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// クレームを取得
	claims, ok := token.Claims.(*jwtConfig)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
