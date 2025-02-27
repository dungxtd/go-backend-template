package bootstrap

import (
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/mongo"
	"github.com/sportgo-app/sportgo-go/postgres"
	"github.com/sportgo-app/sportgo-go/sms"
	"github.com/sportgo-app/sportgo-go/storage"
)

type Application struct {
	Env        *Env
	Mongo      mongo.Client
	Postgres   postgres.Client
	Mailer     email.MailClient
	SmsAdapter sms.SmsAdapter
	Storage    storage.MinioClient
}

// NewDatabase initializes database connections
func NewDatabase(env *Env) (postgres.Client, mongo.Client, error) {
	postgresDB := NewPostgresDatabase(env)
	mongoDB := NewMongoDatabase(env)
	return postgresDB, mongoDB, nil
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()

	// Initialize databases
	postgresDB, mongoDB, err := NewDatabase(app.Env)
	if err != nil {
		panic(err)
	}
	app.Postgres = postgresDB
	app.Mongo = mongoDB

	// Initialize services
	app.Mailer = NewSMTPMailer(app.Env)
	app.SmsAdapter = NewSmsSpeedAdapter(app.Env)
	// app.Storage = NewStorage(app.Env)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
