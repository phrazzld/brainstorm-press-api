package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type CreatePostInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

type UpdatePostInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Body   string `json:"body"`
}
