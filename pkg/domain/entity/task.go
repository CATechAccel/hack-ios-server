package entity

import (
	"time"
)

type Task struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Description *string
	IsDone      bool
	UserID      string
}
