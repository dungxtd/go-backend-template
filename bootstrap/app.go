package bootstrap

import (
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/mongo"
	"github.com/sportgo-app/sportgo-go/sms"
)

type Application struct {
	Env          *Env
	Mongo        mongo.Client
	Mailer       email.MailClient
	UnimtxClient sms.UnimtxClient
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.Mailer = NewSMTPMailer(app.Env)
	app.UnimtxClient = NewUnimtxClient(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
