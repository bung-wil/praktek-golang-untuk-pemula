package routes

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTagRoutes(router *gin.Engine, db *gorm.DB) {
	tagController := controllers.NewTagController(db)

	tagRoutes := router.Group("/tags")
	{
		tagRoutes.POST("/", tagController.CreateTag)
		tagRoutes.GET("/", tagController.GetTags)
	}
}
