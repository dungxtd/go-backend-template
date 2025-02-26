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
	query := `INSERT INTO ` + tr.table + ` (id, title, user_id) VALUES ($1, $2, $3)`
	_, err := tr.db.Exec(c, query, task.ID, task.Title, task.UserID)
	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID string) ([]domain.Task, error) {
	query := `SELECT id, title, user_id FROM ` + tr.table + ` WHERE user_id = $1`
	rows, err := tr.db.Query(c, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err = rows.Scan(&task.ID, &task.Title, &task.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if tasks == nil {
		return []domain.Task{}, nil
	}
	return tasks, nil
}
