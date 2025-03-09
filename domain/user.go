package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

const (
	UserTable = "users"
)

type User struct {
	ID            string    `json:"id" bun:"id,pk"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
	Name          string    `json:"name" bun:"name"`
	Email         string    `json:"email" bun:"email"`
	PhoneNumber   string    `json:"phone_number" bun:"phone_number"`
	Password      string    `json:"-" bun:"password"`
	GoogleID      string    `json:"google_id" bun:"google_id"`
	FacebookID    string    `json:"facebook_id" bun:"facebook_id"`
	AppleID       string    `json:"apple_id" bun:"apple_id"`
	bun.BaseModel `bun:"table:users,alias:u"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	GetByPhone(c context.Context, phoneNumber string) (User, error)
	Update(c context.Context, user *User) error
}
