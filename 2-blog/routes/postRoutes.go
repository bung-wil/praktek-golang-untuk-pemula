package routes

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupPostRoutes(router *gin.Engine, db *gorm.DB) {
	postController := controllers.NewPostController(db)

	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", postController.CreatePost)
		postRoutes.GET("/", postController.GetPosts)
		postRoutes.GET("/:id", postController.GetPost)
		postRoutes.PUT("/:id", postController.UpdatePost)
		postRoutes.DELETE("/:id", postController.DeletePost)
	}
}
