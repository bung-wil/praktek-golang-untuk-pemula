package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes mengatur semua rute untuk aplikasi
func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Setup routes
	SetupContactRoutes(r, db)

	return r
}
