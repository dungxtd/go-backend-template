package migrations

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql
var sqlMigrations embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}

// Migrate runs all migrations
func Migrate() {
	env := bootstrap.NewEnv()

	postgresDB, _, err := bootstrap.NewDatabase(env)
	if err != nil {
		log.Fatal(err)
	}

	db := postgresDB.Database()
	migrator := db.NewMigrator(Migrations)

	err = migrator.Init(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	group, err := migrator.Migrate(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if group.ID == 0 {
		fmt.Printf("there are no new migrations to run\n")
	} else {
		fmt.Printf("migrated to %s\n", group)
	}
}
