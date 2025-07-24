package models

import (
	"time"

	"gorm.io/gorm"
)

// Contact merepresentasikan model data kontak
type Contact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name" validate:"required,min=3"`
	Email     string         `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Phone     string         `gorm:"size:20;not null" json:"phone" validate:"required"`
	Address   string         `gorm:"type:text" json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName mengembalikan nama tabel yang digunakan untuk model Contact
func (Contact) TableName() string {
	return "contacts"
}
