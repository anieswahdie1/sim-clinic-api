package repository

import "sim-clinic-api/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

type RoleRepository interface {
	FindByID(id uint) (*model.Role, error)
}
