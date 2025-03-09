package resutil

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/sportgo-app/sportgo-go/domain"
)

// NewErrorResponse creates an error response based on the current environment
func NewErrorResponse(err error, userMessage ...string) domain.ErrorResponse {
	appEnv := viper.GetString("APP_ENV")

	// Handle development environment
	if appEnv == "development" || appEnv == "dev" {
		// Use null coalescing pattern for userMessage
		var message string
		if len(userMessage) > 0 {
			message = userMessage[0]
		}

		return domain.ErrorResponse{
			Message: message,
			Error:   err.Error(),
		}
	}

	// For production environment
	return domain.ErrorResponse{
		Message: "something went wrong",
	}
}

// HandleErrorResponse sends an error response through the gin context
func HandleErrorResponse(c *gin.Context, statusCode int, err error, userMessage ...string) {
	c.JSON(statusCode, NewErrorResponse(err, userMessage...))
}

// NewSuccessResponse creates a success response with a message
func NewSuccessResponse(message string) domain.SuccessResponse {
	return domain.SuccessResponse{
		Message: message,
	}
}

// HandleSuccessResponse sends a success response through the gin context
func HandleSuccessResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, NewSuccessResponse(message))
}

// HandleDataResponse sends a data response through the gin context
func HandleDataResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
