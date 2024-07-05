package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	TaskName  string `validate:"required,max=255"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

var validate = validator.New()

func (t *Task) Validate() error {
	return validate.Struct(t)
}
