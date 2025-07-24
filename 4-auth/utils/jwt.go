package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"auth/models"
)

// GenerateToken menghasilkan JWT token untuk user
func GenerateToken(user *models.User, jwtSecret string, jwtExpiration string) (string, error) {
	expiration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// VerifyToken memverifikasi JWT token
func VerifyToken(tokenString string, jwtSecret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}