package main

import (
	"log"
	"os"

	"github.com/adrianoyuji/go-rest-api-template/config"
	"github.com/adrianoyuji/go-rest-api-template/internal/routes"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDatabase()

	r := routes.SetupRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
