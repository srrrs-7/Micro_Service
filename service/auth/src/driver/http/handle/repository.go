package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/srrrs-7/micro_service.git/auth/util/env"
	"gorm.io/gorm"
)

type server struct {
	env    *env.Config
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(env *env.Config, db *gorm.DB) *server {
	return &server{
		env:    env,
		db:     db,
		router: gin.Default(),
	}
}
