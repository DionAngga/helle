package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DNS")
}

func PORT() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}
