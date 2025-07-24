package middlewares

import (
	"auth/config"
	"auth/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware memverifikasi JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authorization header is required")
			ctx.Abort()
			return
		}

		// Format: "Bearer <token>"
		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid authorization format")
			ctx.Abort()
			return
		}

		// Verifikasi token
		config := config.LoadConfig()
		token, err := utils.VerifyToken(tokenString[1], config.JWTSecret)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		// Set claims ke context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("userID", claims["id"])
			ctx.Set("userEmail", claims["email"])
		} else {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid token claims")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
