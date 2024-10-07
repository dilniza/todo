package handler

import (
	"net/http"
	"todo/api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUser retrieves a user by ID
// @Summary Retrieve a user by ID
// @Description Get user details by user ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	userID := c.Param("id")
	if userID == "" {
		handleResponseLog(c, h.Log, "User ID is empty", http.StatusBadRequest, "")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		handleResponseLog(c, h.Log, "Invalid user ID format", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Services.User().GetUser(c.Request.Context(), objectID)
	if err != nil {
		handleErrorResponse(c, err, "GetUser", http.StatusInternalServerError)
		return
	}

	handleResponseLog(c, h.Log, "Success", http.StatusOK, user)
}

// UpdateUser updates user information by ID
// @Summary Update user by ID
// @Description Update user details
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUser true "User update data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	userID := c.Param("id")
	if userID == "" {
		handleResponseLog(c, h.Log, "User ID is empty", http.StatusBadRequest, "")
		return
	}

	updateReq := models.UpdateUser{}
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		handleResponseLog(c, h.Log, "Error binding JSON", http.StatusBadRequest, err.Error())
		return
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		handleResponseLog(c, h.Log, "Invalid user ID format", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != updateReq.ID {
		handleResponseLog(c, h.Log, "Unauthorized: cannot update another user's profile", http.StatusUnauthorized, "")
		return
	}

	updatedUser, err := h.Services.User().UpdateUser(c.Request.Context(), objectID, updateReq)
	if err != nil {
		handleErrorResponse(c, err, "UpdateUser", http.StatusInternalServerError)
		return
	}

	handleResponseLog(c, h.Log, "Success", http.StatusOK, updatedUser)
}

// DeleteUser removes a user by ID
// @Summary Delete user by ID
// @Description Delete a user from the system
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	userID := c.Param("id")
	if userID == "" {
		handleResponseLog(c, h.Log, "User ID is empty", http.StatusBadRequest, "")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		handleResponseLog(c, h.Log, "Invalid user ID format", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != userID {
		handleResponseLog(c, h.Log, "Unauthorized: cannot delete another user's profile", http.StatusUnauthorized, "")
		return
	}

	err = h.Services.User().DeleteUser(c.Request.Context(), objectID)
	if err != nil {
		handleErrorResponse(c, err, "DeleteUser", http.StatusInternalServerError)
		return
	}

	handleResponseLog(c, h.Log, "Success", http.StatusOK, "User deleted successfully")
}

// GetAllUsers retrieves all users with optional pagination and search
// @Summary Retrieve all users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of users per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {array} models.User
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Router /users [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Error parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Error parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	allUsers, err := h.Services.User().GetAllUsers(c.Request.Context(), search, page, limit)
	if err != nil {
		handleErrorResponse(c, err, "GetAllUsers", http.StatusInternalServerError)
		return
	}

	handleResponseLog(c, h.Log, "Success", http.StatusOK, allUsers)
}

// GetUserTaskLists retrieves all task lists for a user
// @Summary Retrieve task lists for a user
// @Description Get all task lists for a specific user
// @Tags Users
// @Produce json
// @Param id path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of task lists per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {array} models.TaskList
// @Failure 400 {object} models.ErrorResponse "Invalid user ID format"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /users/{id}/tasklists [get]
func (h *Handler) GetUserTaskLists(c *gin.Context) {
	authInfo, err := getAuthInfo(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	userID := c.Param("id")
	if userID == "" {
		handleResponseLog(c, h.Log, "User ID is empty", http.StatusBadRequest, "")
		return
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		handleResponseLog(c, h.Log, "Invalid user ID format", http.StatusBadRequest, err.Error())
		return
	}

	if authInfo.UserID != userID {
		handleResponseLog(c, h.Log, "Unauthorized: cannot get another user's task lists", http.StatusUnauthorized, "")
		return
	}

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Error parsing page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponseLog(c, h.Log, "Error parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	allTaskLists, err := h.Services.User().GetUserTaskLists(c.Request.Context(), objectID, search, page, limit)
	if err != nil {
		handleErrorResponse(c, err, "GetUserTaskLists", http.StatusInternalServerError)
		return
	}

	handleResponseLog(c, h.Log, "Success", http.StatusOK, allTaskLists)
}
