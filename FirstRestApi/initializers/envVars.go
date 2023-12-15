package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	dir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current working directory: %v", err)
    }

    log.Printf("Current working directory: %s", dir)

	// Load our Environment Variables
	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
}