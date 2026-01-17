package config

import (
	"fmt"
	"log"
	"os"

	"github.com/adrianoyuji/go-rest-api-template/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automigrate
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	return db
}
