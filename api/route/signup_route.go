package route

import (
	"github.com/sportgo-app/sportgo-go/sms"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/controller"
	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/mongo"
	"github.com/sportgo-app/sportgo-go/repository"
	"github.com/sportgo-app/sportgo-go/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, mailer email.MailClient, unimtxClient sms.UnimtxClient, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
		Mailer:        mailer,
		UnimtxClient:  unimtxClient,
	}
	group.POST("/signup", sc.Signup)
}
