package route

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/api/controller"
	"github.com/sportgo-app/sportgo-go/bootstrap"
	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/postgres"
	"github.com/sportgo-app/sportgo-go/repository"
	"github.com/sportgo-app/sportgo-go/usecase"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db postgres.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserTable)
	lc := &controller.LoginController{
		LoginUsecase:      usecase.NewLoginUsecase(ur, timeout),
		SocialAuthUsecase: usecase.NewSocialAuthUsecase(ur, timeout, env),
		Env:               env,
	}

	group.POST("/login/email", lc.LoginWithEmail)
	group.POST("/login/phone", lc.LoginWithPhone)
	group.POST("/login/social", lc.SocialLogin)
	group.GET("/check", lc.CheckUserExists)
}
