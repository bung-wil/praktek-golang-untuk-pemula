package controllers

import (
	"auth/config"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthController struct untuk menampung DB connection
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController menginisialisasi controller baru
func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// Register handler untuk registrasi user
func (ac *AuthController) Register(ctx *gin.Context) {
	var payload *models.RegisterRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Cek apakah email sudah terdaftar
	var existingUser models.User
	result := ac.DB.Where("email = ?", payload.Email).First(&existingUser)
	if result.Error == nil {
		utils.ErrorResponse(ctx, http.StatusConflict, "Email already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Buat user baru
	newUser := models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	}

	result = ac.DB.Create(&newUser)
	if result.Error != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Generate token
	config := config.LoadConfig()
	token, err := utils.GenerateToken(&newUser, config.JWTSecret, config.JWTExp)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Response
	response := models.AuthResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
		Token: token,
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "User registered successfully", response)
}

// Login handler untuk login user
func (ac *AuthController) Login(ctx *gin.Context) {
	var payload *models.LoginRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Cari user berdasarkan email
	var user models.User
	result := ac.DB.Where("email = ?", payload.Email).First(&user)
	if result.Error != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate token
	config := config.LoadConfig()
	token, err := utils.GenerateToken(&user, config.JWTSecret, config.JWTExp)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Response
	response := models.AuthResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Login successful", response)
}

// Logout handler untuk logout user
func (ac *AuthController) Logout(ctx *gin.Context) {
	// Pada implementasi JWT stateless, logout biasanya dilakukan di client side
	// dengan menghapus token dari penyimpanan client
	utils.SuccessResponse(ctx, http.StatusOK, "Logout successful", nil)
}
