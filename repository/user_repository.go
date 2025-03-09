package repository

import (
	"context"
	"time"

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
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	_, err := ur.db.NewInsert().Model(user).Exec(c)
	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	var users []domain.User
	err := ur.db.NewSelect().Model(&users).Scan(c)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return []domain.User{}, nil
	}
	return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.db.NewSelect().Model(&user).Where("email = ?", email).Scan(c)
	return user, err
}

func (ur *userRepository) Update(c context.Context, user *domain.User) error {
	user.UpdatedAt = time.Now()
	_, err := ur.db.NewUpdate().Model(user).WherePK().Exec(c)
	return err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	err := ur.db.NewSelect().Model(&user).Where("id = ?", id).Scan(c)
	return user, err
}

func (ur *userRepository) GetByPhone(c context.Context, phoneNumber string) (domain.User, error) {
	var user domain.User
	err := ur.db.NewSelect().Model(&user).Where("phone_number = ?", phoneNumber).Scan(c)
	return user, err
}
