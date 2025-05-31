package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found.")
	}
}
