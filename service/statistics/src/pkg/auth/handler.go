package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken(ctx *gin.Context) {
	var req genTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// トークンを生成する
	token, err := generateToken(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func VerifyToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bearer toke not found"})
		return
	}
	// "Bearer "を削除してトークンを取得
	token = token[len("Bearer "):]

	// トークンを検証する
	claims, err := verifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"claims": claims,
	})
}
