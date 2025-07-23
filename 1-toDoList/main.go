package main

import (
	"log"
	"os"
	"to-do-list/config"
	"to-do-list/models"
	"to-do-list/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	dbConfig := config.LoadDBConfig()
	config.ConnectDB(dbConfig)
	db := config.GetDB()

	// Auto migrate models
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup routes
	router := routes.SetupRoutes(db)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
