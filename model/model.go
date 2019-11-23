package model

import "github.com/jinzhu/gorm"

// Article required Title, Text
type Article struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}
