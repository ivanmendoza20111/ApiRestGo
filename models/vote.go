package models

import "github.com/jinzhu/gorm"

type Vote struct {
	gorm.Model
	CommentID uint `json:"comment_id" gorm:"not null"`
	UserID    uint `json:"user_id" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
