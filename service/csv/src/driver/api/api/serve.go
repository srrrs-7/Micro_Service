package api

import (
	"github.com/gin-gonic/gin"
	"github.com/srrrs-7/micro_service.git/csv/driver/api/api/middleware"
	"github.com/srrrs-7/micro_service.git/csv/pkg/csv"
	"github.com/srrrs-7/micro_service.git/csv/util/logger"
)

func NewServer() {
	logger.NewLogger()

	router := gin.Default()
	// cors setting
	middleware.NewCors(router)

	router.GET("/csv", csv.CreateCsvHandler)

	router.Run(":8080")
}
