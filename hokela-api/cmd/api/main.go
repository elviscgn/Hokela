package main

import (
	"hokela-api/internal/handlers"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Failed to load .env")
	// }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Println("Database connection successfully established")

	db.AutoMigrate(&handlers.Customer{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := handlers.SetupRouter(db)

	log.Printf("Starting :%s", port)
	r.Run(":" + port)
}
