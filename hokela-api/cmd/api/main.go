package main

import (
	"hokela-api/internal/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Failed to load .env")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", handlers.HealthCheck)

	}

	log.Printf("Starting :%s", port)
	r.Run(":" + port)
}
