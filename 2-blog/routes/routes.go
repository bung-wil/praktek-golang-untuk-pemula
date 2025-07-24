package routes

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Setup routes
	SetupPostRoutes(router, db)
	SetupCategoryRoutes(router, db)
	SetupTagRoutes(router, db)

	return router
}
