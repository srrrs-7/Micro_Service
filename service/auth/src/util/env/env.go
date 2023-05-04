package env

import (
	"os"
)

type Config struct {
	ClientSecret string
	Database     string
	Host         string
	Port         string
}

func LoadEnv() *Config {
	return &Config{
		ClientSecret: setEnv("CLIENT_SECRET", "client_secret"),
		Database:     setEnv("DATABASE", "mysql"),
		Host:         setEnv("HOST", "localhost"),
		Port:         setEnv("PORT", "3306"),
	}
}

func setEnv(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}
	return os.Getenv(key)
}
