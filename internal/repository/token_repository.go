package repository

import (
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
	"time"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) BlacklistToken(token *model.BlacklistedToken) error {
	return r.db.Create(token).Error
}

func (r *tokenRepository) IsTokenBlacklisted(token string) (bool, error) {
	var blacklistedToken model.BlacklistedToken
	result := r.db.Where("token = ? AND expires_at > ?", token, time.Now()).First(&blacklistedToken)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (r *tokenRepository) CleanExpiredTokens() error {
	return r.db.Where("expires_at <= ?", time.Now()).Delete(&model.BlacklistedToken{}).Error
}

func (r *tokenRepository) GetUserActiveTokens(userID uint) ([]model.BlacklistedToken, error) {
	var tokens []model.BlacklistedToken
	err := r.db.Where("user_id = ? AND expires_at > ?", userID, time.Now()).Find(&tokens).Error
	return tokens, err
}
