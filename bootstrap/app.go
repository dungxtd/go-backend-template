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

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	app.Mongo = NewMongoDatabase(app.Env)
	app.Mailer = NewSMTPMailer(app.Env)
	// app.UnimtxClient = NewUnimtxClient(app.Env)
	app.SmsAdapter = NewSmsSpeedAdapter(app.Env)
	//app.Storage = NewStorage(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
