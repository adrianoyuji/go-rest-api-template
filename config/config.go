package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
	// Basic required checks
	required := []string{"DB_URL", "JWT_SECRET"}
	for _, k := range required {
		if os.Getenv(k) == "" {
			log.Printf("Warning: %s is not set\n", k)
		}
	}
}
