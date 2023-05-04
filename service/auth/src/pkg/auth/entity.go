package auth

import "github.com/dgrijalva/jwt-go"

// カスタムのクレームを定義
type jwtConfig struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type genTokenRequest struct {
	UserID string `json:"userId"`
}
