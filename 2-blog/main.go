package main

import (
	"blog/config"
	"blog/models"
	"blog/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	dbConfig := config.LoadDBConfig()
	config.ConnectDB(dbConfig)
	db := config.GetDB()

	// Auto migrate models
	if err := db.AutoMigrate(
		&models.Post{},
		&models.Category{},
		&models.Tag{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup routes
	router := routes.SetupRoutes(db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
