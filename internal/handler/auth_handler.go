package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/service"
	"strings"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, password and role
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "Register Request"
// @Success 201 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 409 {object} map[string]interface{} "User already exists"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var request model.RegisterRequest

	if err := c.Bind(&request); err != nil {
		logrus.Warnf("Invalid register request: %v", err)
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		logrus.Warnf("Validation failed: %v", err)
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	user, err := h.authService.Register(request)
	if err != nil {
		return handleServiceError(c, err)
	}

	logrus.Infof("User registered: %s", user.Username)
	return c.JSON(http.StatusCreated, successResponse(map[string]interface{}{
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role.Name,
		},
	}))
}

// Login godoc
// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login Request"
// @Success 200 {object} map[string]interface{} "Login successful"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var request model.LoginRequest

	if err := c.Bind(&request); err != nil {
		logrus.Warnf("Invalid login request: %v", err)
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		logrus.Warnf("Validation failed: %v", err)
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	response, err := h.authService.Login(request)
	if err != nil {
		return handleServiceError(c, err)
	}

	logrus.Infof("User logged in: %s", request.Username)
	return c.JSON(http.StatusOK, successResponse(response))
}

// Logout godoc
// @Summary Logout user
// @Description Logout user and invalidate JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Logout successful"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c echo.Context) error {
	// Get token from header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, errorResponse("Authorization header required"))
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid authorization format"))
	}

	tokenString := parts[1]

	// Get user ID from context (set oleh auth middleware)
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	// Call logout service
	err := h.authService.Logout(tokenString, userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	logrus.Infof("User %d logged out successfully", userID)
	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "Logout successful",
	}))
}

func successResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"data":    data,
	}
}

func errorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"error":   message,
	}
}

func handleServiceError(c echo.Context, err error) error {
	if serviceErr, ok := err.(*service.ServiceError); ok {
		return c.JSON(serviceErr.Code, errorResponse(serviceErr.Message))
	}
	logrus.Errorf("Internal server error: %v", err)
	return c.JSON(http.StatusInternalServerError, errorResponse("Internal server error"))
}
