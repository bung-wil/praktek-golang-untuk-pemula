package controllers

import (
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

// // CreatePost membuat post baru
func (pc *PostController) CreatePost(c *gin.Context) {
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}

	if err := pc.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": post})
}

// GetPosts mendapatkan semua post
func (pc *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := pc.DB.Preload("Categories").Preload("Tags").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// // GetPost mendapatkan post berdasarkan ID
func (pc *PostController) GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.DB.Preload("Categories").Preload("Tags").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost memperbarui post
func (pc *PostController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// // DeletePost menghapus post
func (pc *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := pc.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	if err := pc.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
