package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL_KEY = "DATABASE_URL"
	JWT_SECRET_KEY   = "JWT_SECRET"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)
}
