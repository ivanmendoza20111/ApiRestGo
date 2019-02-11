package models

import "github.com/jinzhu/gorm"

// Vote permite controlar que  un usuario solo
// vote una unica vez por cada comentario
type Vote struct {
	gorm.Model
	CommentID uint `json:"comment_id" gorm:"not null"`
	UserID    uint `json:"user_id" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
