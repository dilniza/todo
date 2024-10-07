package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"todo/api/models"
	"todo/service"

	"github.com/gin-gonic/gin"
)

// Handler structure
type Handler struct {
	Log      *log.Logger
	Services *service.Service
}

// NewHandler initializes a new handler
func NewHandler(services *service.Service, log *log.Logger) *Handler {
	return &Handler{
		Log:      log,
		Services: services,
	}
}

// handleResponseLog is a utility function to log and send responses
func handleResponseLog(c *gin.Context, logger *log.Logger, message string, statusCode int, data interface{}) {
	if statusCode >= 400 {
		logger.Printf("Error: %s, Status Code: %d, Data: %v", message, statusCode, data)
	}
	c.JSON(statusCode, data)
}

// handleErrorResponse is a utility function to handle error responses
func handleErrorResponse(c *gin.Context, err error, context string, statusCode int) {
	log.Printf("Error in %s: %v", context, err)
	var errorResponse models.ErrorResponse

	// Customize error responses based on the status code
	if statusCode == http.StatusBadRequest {
		errorResponse = models.ErrorResponse{Message: "Invalid request"}
	} else if statusCode == http.StatusUnauthorized {
		errorResponse = models.ErrorResponse{Message: "Unauthorized"}
	} else {
		errorResponse = models.ErrorResponse{Message: "Internal server error"}
	}

	c.JSON(statusCode, errorResponse)
}

// getAuthInfo retrieves authentication information from the context
func getAuthInfo(c *gin.Context) (*models.AuthInfo, error) {
	authInfo, exists := c.Get("authInfo")
	if !exists {
		return nil, errors.New("authentication information not found")
	}
	return authInfo.(*models.AuthInfo), nil
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil || page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "2"
	}

	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil || limit == 0 {
		return 2, nil
	}
	return limit, nil
}
