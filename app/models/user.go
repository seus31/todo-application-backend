package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-" validate:"required,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:,"`
}
