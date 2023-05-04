package main

import (
	"github.com/srrrs-7/micro_service.git/src/driver/api"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// // load config
	// config, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatalf("cannot load config: %v", err)
	// }
	// domain instance
	// start server
	api.NewApiServe()
}
