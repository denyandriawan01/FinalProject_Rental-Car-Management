package initializer

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var (
	DB_CONN 	string
	APP_PORT 	string
)

func EnvLoader() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	// "username:password@tcp(host:port)/dbname"
	DB_CONN  = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	// "http://app_host:app_port"
	APP_PORT = fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
}