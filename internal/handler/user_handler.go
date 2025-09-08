package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"sim-clinic-api/internal/service"
	"strconv"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users with role-based access control
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Users retrieved successfully"
// @Failure 403 {object} map[string]interface{} "Access denied"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users [get]
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	// Get user role from context (set by auth middleware)
	userRole, ok := c.Get("userRole").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	users, err := h.userService.GetAllUsers(userRole)
	if err != nil {
		return handleServiceError(c, err)
	}

	// Filter sensitive data sebelum return
	filteredUsers := make([]map[string]interface{}, len(users))
	for i, user := range users {
		filteredUsers[i] = map[string]interface{}{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role.Name,
			"created_at": user.CreatedAt,
		}
	}

	logrus.Infof("User with role %s retrieved %d users", userRole, len(users))
	return c.JSON(http.StatusOK, successResponse(filteredUsers))
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID with role-based access control
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User retrieved successfully"
// @Failure 403 {object} map[string]interface{} "Access denied"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c echo.Context) error {
	// Get user role from context
	userRole, ok := c.Get("userRole").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	// Get user ID from path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	user, err := h.userService.GetUserByID(uint(id), userRole)
	if err != nil {
		return handleServiceError(c, err)
	}

	// Filter sensitive data
	filteredUser := map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role.Name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	logrus.Infof("User with role %s retrieved user ID %d", userRole, user.ID)
	return c.JSON(http.StatusOK, successResponse(filteredUser))
}
