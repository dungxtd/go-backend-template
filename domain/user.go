package domain

import (
	"context"

	"github.com/uptrace/bun"
)

const (
	UserTable = "users"
)

type User struct {
	ID            string `json:"id" bun:"id,pk"`
	Name          string `json:"name" bun:"name"`
	Email         string `json:"email" bun:"email"`
	PhoneNumber   string `json:"phone_number" bun:"phone_number"`
	Password      string `json:"-" bun:"password"`
	bun.BaseModel `bun:"table:users,alias:u"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	GetByPhone(c context.Context, phoneNumber string) (User, error)
}
