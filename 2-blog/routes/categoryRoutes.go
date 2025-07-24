package routes

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	categoryController := controllers.NewCategoryController(db)

	categoryRoutes := router.Group("/categories")
	{
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.GET("/", categoryController.GetCategories)
	}
}
