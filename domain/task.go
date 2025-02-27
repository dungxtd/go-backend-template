package domain

import (
	"context"
)

const (
	TaskTable = "tasks"
)

type Task struct {
	ID     string `json:"-" db:"id"`
	Title  string `form:"title" binding:"required" json:"title" db:"title"`
	UserID string `json:"-" db:"user_id"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
