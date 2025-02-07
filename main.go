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

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, app.Mailer, app.SmsAdapter, gin)

	gin.Run(env.ServerAddress)
}
