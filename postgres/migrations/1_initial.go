package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

const usersTable = `
CREATE TABLE users (
	id uuid NOT NULL,
	name text NOT NULL,
	email text,
	password text NOT NULL,
	phone_number text,
	role SMALLINT CHECK (role BETWEEN 0 AND 2) DEFAULT 0, -- 0: customer, 1: owner, 2: admin,
	created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
	updated_at timestamp with time zone DEFAULT current_timestamp,
	PRIMARY KEY (id)
)`

const venuesTable = `
CREATE TABLE venues (
    id SERIAL PRIMARY KEY,
    owner_id INT REFERENCES users(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    latitude DECIMAL(9,6) NOT NULL,  -- Vĩ độ
    longitude DECIMAL(9,6) NOT NULL, -- Kinh độ
    sport_type SMALLINT CHECK (sport_type BETWEEN 0 AND 3), -- 0: football, 1: basketball, 2: tennis, 3: badminton
    price_per_hour DECIMAL(10,2) CHECK (price_per_hour >= 0),
    created_at TIMESTAMP DEFAULT NOW()
);`

const bookingsTable = `
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    venue_id INT REFERENCES venues(id) ON DELETE CASCADE,
    booking_date DATE NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL CHECK (end_time > start_time),
    status SMALLINT CHECK (status BETWEEN 0 AND 2) DEFAULT 0, -- 0: pending, 1: confirmed, 2: canceled
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (venue_id, booking_date, start_time) -- Ensure no overlapping bookings
);`

const savedVenuesTable = `
CREATE TABLE saved_venues (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    venue_id INT REFERENCES venues(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (user_id, venue_id)
);`

func init() {
	up := []string{
		usersTable,
		venuesTable,
		savedVenuesTable,
		bookingsTable,
	}

	down := []string{
		`DROP TABLE users`,
		`DROP TABLE venues`,
		`DROP TABLE bookings`,
		`DROP TABLE saved_venues`,
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
