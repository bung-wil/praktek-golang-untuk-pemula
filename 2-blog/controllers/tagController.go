package controllers

import (
	"net/http"

	"blog/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagController struct {
	DB *gorm.DB
}

func NewTagController(db *gorm.DB) *TagController {
	return &TagController{DB: db}
}

// CreateTag membuat tag baru
func (tc *TagController) CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// GetTags mendapatkan semua tag
func (tc *TagController) GetTags(c *gin.Context) {
	var tags []models.Tag
	if err := tc.DB.Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tags"})
		return
	}

	c.JSON(http.StatusOK, tags)
}
