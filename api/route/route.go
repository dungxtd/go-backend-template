package route

import (
	"github.com/sportgo-app/sportgo-go/sms"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/middleware"
	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, mailer email.MailClient, unimtxClient sms.UnimtxClient, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, mailer, unimtxClient, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}
