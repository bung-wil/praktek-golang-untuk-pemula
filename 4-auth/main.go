package main

import (
	"auth/config"
	"auth/models"
	"auth/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load konfigurasi
	configuration := config.LoadConfig()

	// Inisialisasi database
	db, err := config.InitDB(configuration)
	if err != nil {
		log.Fatal("Failed to initialize database")
	}

	// Auto migrate model User
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}

	// Inisialisasi Gin
	router := gin.Default()

	// Setup routes
	routes.AuthRoutes(router, db)

	// Jalankan server
	log.Println("Server running on port 8080")
	router.Run(":8080")
}
