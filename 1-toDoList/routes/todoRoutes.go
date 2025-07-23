package routes

import (
	"to-do-list/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes mengatur semua rute untuk aplikasi
func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// Setup routes
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.POST("/", controllers.NewTodoController(db).CreateTodo)
		todoRoutes.GET("/", controllers.NewTodoController(db).GetTodos)
		todoRoutes.GET("/:id", controllers.NewTodoController(db).GetTodo)
		todoRoutes.PUT("/:id", controllers.NewTodoController(db).UpdateTodo)
		todoRoutes.DELETE("/:id", controllers.NewTodoController(db).DeleteTodo)
	}
	return router
}
