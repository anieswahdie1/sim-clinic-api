package service

import "sim-clinic-api/internal/model"

type AuthService interface {
	Register(request model.RegisterRequest) (*model.User, error)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
}
