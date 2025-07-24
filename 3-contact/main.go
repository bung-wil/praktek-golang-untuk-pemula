package main

import (
	"contact/config"
	"contact/models"
	"contact/routes"
	"log"

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
	db.AutoMigrate(models.Contact{})

	// Setup router
	router := routes.SetupRoutes(db)

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
