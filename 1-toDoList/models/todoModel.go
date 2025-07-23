package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo merepresentasikan model untuk item to-do
type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Completed   bool           `gorm:"default:false" json:"completed"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName mengembalikan nama tabel untuk model Todo
func (Todo) TableName() string {
	return "todos"
}
