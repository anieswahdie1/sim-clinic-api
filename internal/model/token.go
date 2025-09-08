package model

import (
	"gorm.io/gorm"
	"time"
)

type BlacklistedToken struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Token     string         `json:"token" gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time      `json:"expires_at" gorm:"not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Reason    string         `json:"reason" gorm:"default:'logout'"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
