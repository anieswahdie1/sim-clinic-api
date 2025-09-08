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

type TokenRepository interface {
	BlacklistToken(token *model.BlacklistedToken) error
	IsTokenBlacklisted(token string) (bool, error)
	CleanExpiredTokens() error
	GetUserActiveTokens(userID uint) ([]model.BlacklistedToken, error)
}
