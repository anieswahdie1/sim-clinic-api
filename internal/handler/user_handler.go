package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"sim-clinic-api/internal/model"
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

// UpdateUser godoc
// @Summary Update user
// @Description Update user with role-based access control
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body model.UpdateUserRequest true "Update Request"
// @Success 200 {object} map[string]interface{} "User updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 403 {object} map[string]interface{} "Access denied"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c echo.Context) error {
	// Get current user info from context
	userRole, ok := c.Get("userRole").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	// Get target user ID from path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	var request model.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	user, err := h.userService.UpdateUser(uint(id), request, userRole, userID)
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

	logrus.Infof("User %d updated user %d successfully", userID, id)
	return c.JSON(http.StatusOK, successResponse(filteredUser))
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user with role-based access control
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted successfully"
// @Failure 403 {object} map[string]interface{} "Access denied"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c echo.Context) error {
	// Get current user info from context
	userRole, ok := c.Get("userRole").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, errorResponse("Invalid user context"))
	}

	// Get target user ID from path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid user ID"))
	}

	err = h.userService.DeleteUser(uint(id), userRole, userID)
	if err != nil {
		return handleServiceError(c, err)
	}

	logrus.Infof("User %d deleted user %d successfully", userID, id)
	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "User deleted successfully",
	}))
}
