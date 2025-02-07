package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type TaskStatus string

const (
	NotYetStarted TaskStatus = "not_yet_started"
	InProgress    TaskStatus = "in_progress"
	Completed     TaskStatus = "completed"
)

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
)

type Task struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	TaskName  string     `json:"task_name" validate:"required,max=255"`
	UserID    uint       `json:"user_id" gorm:"index"`
	ParentID  *uint      `json:"parent_id" gorm:"index"`
	Parent    *Task      `json:"parent" gorm:"foreignKey:ParentID" validate:"-"`
	Children  []Task     `json:"children" gorm:"foreignKey:id" validate:"-"`
	DueDate   *string    `json:"due_date" gorm:"type:date" validate:"omitempty"`
	DueTime   *string    `json:"due_time" gorm:"type:time" validate:"omitempty"`
	Status    TaskStatus `json:"status" gorm:"type:task_status" validate:"required,valid_task_status"`
	Priority  Priority   `json:"priority" gorm:"type:int" validate:"required,valid_priority"`
	CreatedAt time.Time  `json:"created_at" validate:"-"`
	UpdatedAt time.Time  `json:"updated_at" validate:"-"`
}

type TaskOption func(*Task)

func WithParentID(parentID *uint) TaskOption {
	return func(t *Task) {
		if parentID != nil {
			t.ParentID = parentID
		}
	}
}

func WithDueDate(dueDate *string) TaskOption {
	return func(t *Task) {
		if dueDate != nil {
			t.DueDate = dueDate
		}
	}
}

func WithDueTime(dueTime *string) TaskOption {
	return func(t *Task) {
		if dueTime != nil {
			t.DueTime = dueTime
		}
	}
}

func WithStatus(status TaskStatus) TaskOption {
	return func(t *Task) {
		if status != "" {
			t.Status = status
		}
	}
}

func WithPriority(priority Priority) TaskOption {
	return func(t *Task) {
		if priority != 0 {
			t.Priority = priority
		}
	}
}

func NewTask(taskName string, userID uint, options ...TaskOption) *Task {
	task := &Task{
		TaskName: taskName,
		UserID:   userID,
	}

	for _, opt := range options {
		opt(task)
	}

	return task
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("valid_task_status", isValidTaskStatus)
	_ = validate.RegisterValidation("valid_priority", isValidPriority)
}

func GetTaskValidator() *validator.Validate {
	return validate
}

func isValidTaskStatus(fl validator.FieldLevel) bool {
	status := TaskStatus(fl.Field().String())
	return status == NotYetStarted || status == InProgress || status == Completed
}

func isValidPriority(fl validator.FieldLevel) bool {
	priority, ok := fl.Field().Interface().(Priority)
	if !ok {
		return false
	}
	return priority >= Low && priority <= High
}

func (t *Task) BeforeSave(tx *gorm.DB) error {
	if err := validate.Struct(t); err != nil {
		return err
	}
	return nil
}
