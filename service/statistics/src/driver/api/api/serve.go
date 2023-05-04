package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/srrrs-7/micro_service.git/statistics/driver/api/api/middleware"
	"github.com/srrrs-7/micro_service.git/statistics/pkg/auth"
	"github.com/srrrs-7/micro_service.git/statistics/util/logger"
)

func NewServer() {
	logger.NewLogger()

	router := gin.Default()
	// cors setting
	middleware.NewCors(router)

	router.GET("/auth/generate", auth.GenerateToken)
	router.GET("/auth/verify", auth.VerifyToken)

	router.Run(":8888")
	log.Println("http server started :8888")
}
