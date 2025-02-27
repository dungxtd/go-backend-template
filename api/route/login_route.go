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
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
