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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db postgres.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserTable)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
