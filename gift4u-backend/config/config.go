package config

import (
	"fmt"
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
		fmt.Println("Error loading .env file")
		// log.Fatalf("Error loading .env file")
	}

	// log.Println(os.Getenv("DATABASEURI"))
	// log.Println(config)
	config.DATABASEURI = os.Getenv("DATABASEURI")
}

func GetConfig() Config {
	return config
}
