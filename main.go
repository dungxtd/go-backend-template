package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/route"
	"github.com/sportgo-app/sportgo-go/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	// mongo := app.Mongo.Database(env.MongoDBName)
	defer app.CloseDBConnection()

	postgres := app.Postgres.Database()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, postgres, app.Mailer, app.SmsAdapter, gin)

	gin.Run(env.ServerAddress)
}
