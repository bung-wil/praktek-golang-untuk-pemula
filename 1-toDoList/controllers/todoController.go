package controllers

import (
	"net/http"

	"to-do-list/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoController adalah controller untuk operasi CRUD todo
type TodoController struct {
	DB *gorm.DB
}

// NewTodoController membuat instance baru TodoController
func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{DB: db}
}

// CreateTodo membuat item todo baru
func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodos mengambil semua item todo
func (tc *TodoController) GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := tc.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTodo mengambil satu item todo berdasarkan ID
func (tc *TodoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo memperbarui item todo berdasarkan ID
func (tc *TodoController) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo menghapus item todo berdasarkan ID
func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := tc.DB.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todo"})
		return
	}

	if err := tc.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
