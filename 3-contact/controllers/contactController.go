package controllers

import (
	"net/http"

	"contact/models"
	"contact/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ContactController struct untuk menangani operasi terkait kontak
type ContactController struct {
	DB *gorm.DB
}

// NewContactController membuat instance baru dari ContactController
func NewContactController(db *gorm.DB) *ContactController {
	return &ContactController{DB: db}
}

// CreateContact menangani pembuatan kontak baru
func (cc *ContactController) CreateContact(c *gin.Context) {
	var contact models.Contact

	// Bind JSON request ke struct Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input
	if err := utils.ValidateStruct(contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan ke database
	if err := cc.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contact"})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

// GetContacts menangani pengambilan semua kontak
func (cc *ContactController) GetContacts(c *gin.Context) {
	var contacts []models.Contact

	if err := cc.DB.Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

// GetContact menangani pengambilan satu kontak berdasarkan ID
func (cc *ContactController) GetContact(c *gin.Context) {
	id := c.Param("id")
	var contact models.Contact

	if err := cc.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// UpdateContact menangani pembaruan kontak
func (cc *ContactController) UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var contact models.Contact

	// Cari kontak berdasarkan ID
	if err := cc.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input
	if err := utils.ValidateStruct(contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan
	if err := cc.DB.Save(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

// DeleteContact menangani penghapusan kontak
func (cc *ContactController) DeleteContact(c *gin.Context) {
	id := c.Param("id")
	var contact models.Contact

	// Cari kontak berdasarkan ID
	if err := cc.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	// Hapus kontak
	if err := cc.DB.Delete(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete contact"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
