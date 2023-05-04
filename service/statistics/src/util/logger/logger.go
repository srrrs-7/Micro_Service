package logger

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func NewLogger() {
	// Loggerの初期化
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
