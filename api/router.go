package api

import (
	"errors"
	"fmt"
	"net/http"
	"todo/api/handler"
	"todo/pkg/logger"
	"todo/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "todo/api/docs"
)

// Server structure
type Server struct {
	Router  *gin.Engine
	Handler handler.Handler
	Log     logger.ILogger
}

// New initializes a new API server
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server for a todo application.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services *service.Service, log logger.ILogger) Server {
	router := gin.Default()

	h := handler.NewHandler(services, log)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api")
	{
		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/login", h.LoginUser)
			authGroup.POST("/register", h.UserRegister)
			authGroup.POST("/register-confirm", h.UserRegisterConfirm)
			authGroup.PATCH("/user", h.ChangePasswordUser)
		}
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/:id", h.GetUser)
			userGroup.PATCH("/:id", h.UpdateUser)
			userGroup.DELETE("/:id", h.DeleteUser)
			userGroup.GET("", h.GetAllUsers)
			userGroup.GET("/:id/task-lists", h.GetUserTaskLists)
		}

		taskGroup := apiGroup.Group("/task")
		{
			taskGroup.POST("", h.CreateTask)
			taskGroup.GET("/:id", h.GetTask)
			taskGroup.PATCH("/:id", h.UpdateTask)
			taskGroup.DELETE("/:id", h.DeleteTask)
		}

		taskListGroup := apiGroup.Group("/task-list")
		{
			taskListGroup.POST("", h.CreateTaskList)
			taskListGroup.GET("/:id", h.GetTaskList)
			taskListGroup.PATCH("/:id", h.UpdateTaskList)
			taskListGroup.DELETE("/:id", h.DeleteTaskList)
			taskListGroup.GET("", h.GetAllTaskLists)
		}

		labelGroup := apiGroup.Group("/label")
		{
			labelGroup.POST("", h.CreateLabel)
			labelGroup.GET("/:id", h.GetLabel)
			labelGroup.PATCH("/:id", h.UpdateLabel)
			labelGroup.DELETE("/:id", h.DeleteLabel)
			labelGroup.GET("", h.GetAllLabels)
		}
	}

	return Server{
		Router:  router,
		Handler: h,
		Log:     log,
	}
}

// Run starts the server
func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}

// authMiddleware handles authentication checks
func authMiddleware(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	c.Next()
}

// logMiddleware logs request headers
func logMiddleware(c *gin.Context) {
	headers := c.Request.Header
	for key, values := range headers {
		for _, v := range values {
			fmt.Printf("Header: %v, Value: %v\n", key, v)
		}
	}
	c.Next()
}
