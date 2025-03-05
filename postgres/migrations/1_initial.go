package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

const userTable = `
CREATE TABLE users (
id uuid NOT NULL,
created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
updated_at timestamp with time zone DEFAULT current_timestamp,
name text NOT NULL,
email text NOT NULL UNIQUE,
password text NOT NULL,
phone_number text,
PRIMARY KEY (id)
)`

const taskTable = `
CREATE TABLE tasks (
id uuid NOT NULL,
created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
updated_at timestamp with time zone DEFAULT current_timestamp,
title text NOT NULL,
user_id uuid NOT NULL REFERENCES users(id),
PRIMARY KEY (id)
)`

func init() {
	up := []string{
		userTable,
		taskTable,
	}

	down := []string{
		`DROP TABLE tasks`,
		`DROP TABLE users`,
	}

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
