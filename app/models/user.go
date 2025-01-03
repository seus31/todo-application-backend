package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `validate:"required,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=64,containsany=!@#$%^&*()_+-=[]{};:,"`
}
