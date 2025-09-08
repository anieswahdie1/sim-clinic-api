package service

import "sim-clinic-api/internal/model"

type AuthService interface {
	Register(request model.RegisterRequest) (*model.User, error)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
	Logout(tokenString string, userID uint) error
	ValidateToken(tokenString string) (*model.User, error)
}

type UserService interface {
	GetAllUsers(currentUserRole string) ([]model.User, error)
	GetUserByID(id uint, currentUserRole string) (*model.User, error)
	UpdateUser(id uint, request model.UpdateUserRequest, currentUserRole string, currentUserID uint) (*model.User, error)
	DeleteUser(id uint, currentUserRole string, currentUserID uint) error
}
