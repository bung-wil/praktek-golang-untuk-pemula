package routes

import (
	"contact/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupContactRoutes mengatur semua rute yang terkait dengan kontak
func SetupContactRoutes(r *gin.Engine, db *gorm.DB) {
	contactController := controllers.NewContactController(db)

	contactRoutes := r.Group("/api/contacts")
	{
		contactRoutes.GET("/", contactController.GetContacts)
		contactRoutes.POST("/", contactController.CreateContact)
		contactRoutes.GET("/:id", contactController.GetContact)
		contactRoutes.PUT("/:id", contactController.UpdateContact)
		contactRoutes.DELETE("/:id", contactController.DeleteContact)
	}
}
