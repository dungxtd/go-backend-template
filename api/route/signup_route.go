package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/controller"
	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/email"
	"github.com/sportgo-app/sportgo-go/postgres"
	"github.com/sportgo-app/sportgo-go/repository"
	"github.com/sportgo-app/sportgo-go/sms"
	"github.com/sportgo-app/sportgo-go/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db postgres.Database, mailer email.MailClient, smsAdapter sms.SmsAdapter, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserTable)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
		Mailer:        mailer,
		SmsAdapter:    smsAdapter,
	}
	group.POST("/signup", sc.Signup)
}
