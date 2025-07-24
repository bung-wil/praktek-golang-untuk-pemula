package routes

import (
	"auth/controllers"
	middlewares "auth/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthRoutes mengatur routing untuk autentikasi
func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/logout", middlewares.AuthMiddleware(), authController.Logout)
	}
}
