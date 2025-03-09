package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/sportgo-app/sportgo-go/domain"
	"github.com/sportgo-app/sportgo-go/internal/resutil"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	userID := c.GetString("x-user-id")
	// Generate UUID for task ID
	task.ID = uuid.New().String()

	// Convert string userID to UUID
	task.UserID = userID
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	resutil.HandleSuccessResponse(c, http.StatusOK, "Task created successfully")
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		resutil.HandleErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	resutil.HandleDataResponse(c, http.StatusOK, tasks)
}
