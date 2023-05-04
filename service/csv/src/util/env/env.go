package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientSecret string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load("/Users/srrrs/App/Go/App/micro_service/auth/.env")
	if err != nil {
		log.Fatal("Failed to load .env file: ", err)
	}

	env := os.Getenv("CLIENT_SECRET")

	return &Config{
		ClientSecret: env,
	}, nil
}
