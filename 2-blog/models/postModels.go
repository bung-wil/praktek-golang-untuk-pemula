package models

import "gorm.io/gorm"

// Post model untuk artikel blog
type Post struct {
	gorm.Model
	Title      string      `gorm:"size:200;not null" json:"title"`
	Content    string      `gorm:"type:text;not null" json:"content"`
	Categories []*Category `gorm:"many2many:post_categories;" json:"categories"`
	Tags       []*Tag      `gorm:"many2many:post_tags;" json:"tags"`
}

type CreatePostInput struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CategoryIDs []uint `json:"category_ids"`
	TagIDs      []uint `json:"tag_ids"`
}
