package cmd

import (
	"github.com/spf13/cobra"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/route"
	"github.com/sportgo-app/sportgo-go/bootstrap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		app := bootstrap.App()

		env := app.Env

		// mongo := app.Mongo.Database(env.MongoDBName)
		defer app.CloseDBConnection()

		postgres := app.Postgres.Database()

		timeout := time.Duration(env.ContextTimeout) * time.Second

		gin := gin.Default()

		route.Setup(env, timeout, postgres, app.Mailer, app.SmsAdapter, gin)

		gin.Run(env.ServerAddress)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
