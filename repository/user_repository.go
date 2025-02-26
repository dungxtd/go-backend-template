package repository

import (
	"context"

	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/postgres"
)

type userRepository struct {
	db    postgres.Database
	table string
}

func NewUserRepository(db postgres.Database, table string) domain.UserRepository {
	return &userRepository{
		db:    db,
		table: table,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	query := `INSERT INTO ` + ur.table + ` (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := ur.db.Exec(c, query, user.ID, user.Name, user.Email, user.Password)
	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	query := `SELECT id, name, email FROM ` + ur.table
	rows, err := ur.db.Query(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if users == nil {
		return []domain.User{}, nil
	}
	return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	query := `SELECT id, name, email, password FROM ` + ur.table + ` WHERE email = $1`
	var user domain.User
	err := ur.db.QueryRow(c, query).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	query := `SELECT id, name, email, password FROM ` + ur.table + ` WHERE id = $1`
	var user domain.User
	err := ur.db.QueryRow(c, query).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}
