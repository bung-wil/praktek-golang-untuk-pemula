package models

import "gorm.io/gorm"

// Tag model untuk tag blog
type Tag struct {
	gorm.Model
	Name  string  `gorm:"size:100;not null;unique" json:"name"`
	Posts []*Post `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}
