package main

import (
	"github.com/srrrs-7/micro_service.git/auth/driver/database"

	"github.com/srrrs-7/micro_service.git/auth/util/env"
	"github.com/srrrs-7/micro_service.git/auth/util/logger"
)

func main() {
	logger.NewLogger()

	env := env.LoadEnv()
	db := database.NewDb(env)
	server := server.NewServer(env, db)

	// repository
	//

	server.NewRouter()
}
