package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `validate:"required,max=255"`
}
