package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
