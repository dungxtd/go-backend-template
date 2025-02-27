package repository

import (
	"context"

	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/postgres"
)

type taskRepository struct {
	db    postgres.Database
	table string
}

func NewTaskRepository(db postgres.Database, table string) domain.TaskRepository {
	return &taskRepository{
		db:    db,
		table: table,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	_, err := tr.db.NewInsert().Model(task).Exec(c)
	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	var tasks []domain.Task
	err := tr.db.NewSelect().Model(&tasks).Where("user_id = ?", userID).Scan(c)
	if err != nil {
		return nil, err
	}
	if tasks == nil {
		return []domain.Task{}, nil
	}
	return tasks, nil
}
