package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/internal/resutil"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	resutil.HandleDataResponse(c, http.StatusOK, profile)
}
