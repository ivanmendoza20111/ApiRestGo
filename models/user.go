package models

import "github.com/jinzhu/gorm"

// User usuario del sistema
type User struct {
	gorm.Model
	UserName        string    `json:"user_name" gorm:"not null; unique"`
	Email           string    `json:"email" gorm:"not null; unique"`
	FullName        string    `json:"full_name" gorm:"not null"`
	Password        string    `json:"password, omitempty" gorm:"not null; type:varchar(256)"`
	ConfirmPassword string    `json:"confirm_password, omitempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"comments, omitempty"`
}
