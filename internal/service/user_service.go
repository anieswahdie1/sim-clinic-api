package service

import (
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

func (s *userService) GetAllUsers(currentUserRole string) ([]model.User, error) {
	switch currentUserRole {
	case "super_admin":
		// Super admin bisa melihat semua user
		users, err := s.userRepo.FindAll()
		if err != nil {
			return nil, err
		}
		return users, nil

	case "admin":
		// Admin bisa melihat admin dan user (tidak bisa melihat super_admin)
		users, err := s.userRepo.FindByRoles([]string{"admin", "user"})
		if err != nil {
			return nil, err
		}
		return users, nil

	case "user":
		// User tidak boleh melihat user lain
		return nil, &ServiceError{
			Message: "access denied: insufficient permissions",
			Code:    403,
		}

	default:
		return nil, &ServiceError{
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
