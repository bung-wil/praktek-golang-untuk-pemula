package models

import "gorm.io/gorm"

// Category model untuk kategori blog
type Category struct {
	gorm.Model
	Name  string  `gorm:"size:100;not null;unique" json:"name"`
	Posts []*Post `gorm:"many2many:post_categories;" json:"posts,omitempty"`
}
