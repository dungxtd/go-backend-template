package domain

import (
	"context"
	"time"
	"github.com/uptrace/bun"
)

const (
	TaskTable = "tasks"
)

type Task struct {
	ID            string    `json:"-" bun:"id,pk"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
	Title         string    `form:"title" binding:"required" json:"title" bun:"title"`
	UserID        string    `json:"-" bun:"user_id"`
	bun.BaseModel `bun:"table:tasks,alias:t"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
