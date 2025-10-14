package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAllUsers(currentUserRole string, page, limit, search string) ([]model.User, int64, error) {
	switch currentUserRole {
	case "super_admin":
		// Super admin bisa melihat semua user
		users, total, err := s.userRepo.FindAll(page, limit, search)
		if err != nil {
			return nil, 0, err
		}
		return users, total, nil

	case "admin":
		// Admin bisa melihat admin dan user (tidak bisa melihat super_admin)
		users, err := s.userRepo.FindByRoles([]string{"admin", "user"})
		if err != nil {
			return nil, 0, err
		}
		return users, 0, nil

	case "user":
		// User tidak boleh melihat user lain
		return nil, 0, &ServiceError{
			Message: "access denied: insufficient permissions",
			Code:    403,
		}

	default:
		return nil, 0, &ServiceError{
			Message: "invalid user role",
			Code:    400,
		}
	}
}

func (s *userService) GetUserByID(id uint, currentUserRole string) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "user not found", Code: 404}
		}
		return nil, err
	}

	// Authorization check
	if !s.hasPermission(currentUserRole, user.Role.Name) {
		return nil, &ServiceError{
			Message: "access denied: insufficient permissions",
			Code:    403,
		}
	}

	return user, nil
}

func (s *userService) UpdateUser(id uint, request model.UpdateUserRequest, currentUserRole string, currentUserID uint) (*model.User, error) {
	// Get target user
	targetUser, err := s.userRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &ServiceError{Message: "user not found", Code: 404}
		}
		return nil, err
	}

	// Authorization check
	if !s.canUpdateUser(currentUserRole, currentUserID, targetUser, request.RoleID) {
		return nil, &ServiceError{
			Message: "access denied: insufficient permissions",
			Code:    403,
		}
	}

	// Apply updates
	if request.Username != nil {
		// Check if new username is available (excluding current user)
		existingUser, _ := s.userRepo.FindByUsername(*request.Username)
		if existingUser != nil && existingUser.ID != id {
			return nil, &ServiceError{Message: "username already exists", Code: 400}
		}
		targetUser.Username = *request.Username
	}

	if request.Email != nil {
		// Check if new email is available (excluding current user)
		existingEmail, _ := s.userRepo.FindByEmail(*request.Email)
		if existingEmail != nil && existingEmail.ID != id {
			return nil, &ServiceError{Message: "email already exists", Code: 400}
		}
		targetUser.Email = *request.Email
	}

	if request.RoleID != nil {
		// Validate new role exists
		_, err := s.userRepo.FindByID(*request.RoleID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, &ServiceError{Message: "role not found", Code: 400}
			}
			return nil, err
		}
		targetUser.RoleID = *request.RoleID
	}

	// Save updates
	if err := s.userRepo.Update(targetUser); err != nil {
		return nil, err
	}

	// Reload user with role data
	updatedUser, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	logrus.Infof("User %d updated user %d successfully", currentUserID, id)
	return updatedUser, nil
}

func (s *userService) DeleteUser(id uint, currentUserRole string, currentUserID uint) error {
	// Get target user
	targetUser, err := s.userRepo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &ServiceError{Message: "user not found", Code: 404}
		}
		return err
	}

	// Authorization check
	if !s.canDeleteUser(currentUserRole, currentUserID, targetUser) {
		return &ServiceError{
			Message: "access denied: insufficient permissions",
			Code:    403,
		}
	}

	// Prevent self-deletion
	if currentUserID == id {
		return &ServiceError{Message: "cannot delete yourself", Code: 400}
	}

	if err := s.userRepo.Delete(id); err != nil {
		return err
	}

	logrus.Infof("User %d deleted user %d successfully", currentUserID, id)
	return nil
}

// hasPermission checks if current user role has permission to access target user role
func (s *userService) hasPermission(currentRole, targetRole string) bool {
	roleHierarchy := map[string]int{
		"super_admin": 3,
		"admin":       2,
		"user":        1,
	}

	currentLevel, currentOk := roleHierarchy[currentRole]
	targetLevel, targetOk := roleHierarchy[targetRole]

	if !currentOk || !targetOk {
		return false
	}

	// Current user can only access users with equal or lower level
	return currentLevel >= targetLevel
}

// Helper function to get role hierarchy level
func (s *userService) getRoleLevel(role string) int {
	roleLevels := map[string]int{
		"super_admin": 3,
		"admin":       2,
		"user":        1,
	}
	return roleLevels[role]
}

func (s *userService) canUpdateUser(currentRole string, currentUserID uint, targetUser *model.User, newRoleID *uint) bool {
	roleHierarchy := map[string]int{
		"super_admin": 3,
		"admin":       2,
		"user":        1,
	}

	//currentLevel := roleHierarchy[currentRole]
	targetLevel := roleHierarchy[targetUser.Role.Name]

	// Super admin can update anyone
	if currentRole == "super_admin" {
		return true
	}

	// Admin can update themselves and users
	if currentRole == "admin" {
		// Admin can update themselves
		if currentUserID == targetUser.ID {
			return true
		}
		// Admin can update users (but not other admins or super_admins)
		return targetLevel <= roleHierarchy["user"]
	}

	// User can only update themselves
	if currentRole == "user" {
		return currentUserID == targetUser.ID
	}

	return false
}

func (s *userService) canDeleteUser(currentRole string, currentUserID uint, targetUser *model.User) bool {
	roleHierarchy := map[string]int{
		"super_admin": 3,
		"admin":       2,
		"user":        1,
	}

	//currentLevel := roleHierarchy[currentRole]
	targetLevel := roleHierarchy[targetUser.Role.Name]

	// Super admin can delete anyone except themselves
	if currentRole == "super_admin" {
		return currentUserID != targetUser.ID
	}

	// Admin can only delete users
	if currentRole == "admin" {
		return targetLevel <= roleHierarchy["user"]
	}

	// User cannot delete anyone
	return false
}
