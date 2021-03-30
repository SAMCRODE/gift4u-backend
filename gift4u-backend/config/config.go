package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASEURI string
}

var config Config

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		// fmt.Println("Error loading .env file")
		log.Fatalf("Error loading .env file")
	}

	config.DATABASEURI = os.Getenv("DATABASE_URL")
}

func GetConfig() Config {
	return config
}
