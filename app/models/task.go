package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName string `validate:"required,max=255"`
}
