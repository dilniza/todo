package handler

import (
	"net/http"
	"todo/api/models"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login
// @Description Authenticate user
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User credentials"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var creds models.LoginRequest
	if err := c.ShouldBindJSON(&creds); err != nil {
		handleErrorResponse(c, err, "Login: binding JSON", http.StatusBadRequest)
		return
	}

	// Use the AuthService to authenticate the user
	authResponse, err := h.Services.AuthService.Login(creds.Email, creds.Password)
	if err != nil {
		handleErrorResponse(c, err, "Login: authentication", http.StatusUnauthorized)
		return
	}

	// Store authInfo in context
	authInfo := &models.AuthInfo{
		UserID: authResponse.User.ID,
		Email:  creds.Email,
		Role:   authResponse.User.Role, // If applicable
	}
	c.Set("authInfo", authInfo)

	handleResponseLog(c, h.Log, "Login successful", http.StatusOK, authResponse)
}
