package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null" valid:"required,alphanum,length(3|50)"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null" valid:"required,email"`
	Password  string         `json:"-" gorm:"not null" valid:"required,length(6|100)"`
	RoleID    uint           `json:"role_id" gorm:"not null"`
	Role      Role           `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}

type RegisterRequest struct {
	Username string `json:"username" valid:"required,alphanum,length(3|50)"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,length(6|100)"`
	RoleID   uint   `json:"role_id" valid:"required"`
}

func (r *RegisterRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}

type LoginRequest struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func (l *LoginRequest) Validate() error {
	_, err := govalidator.ValidateStruct(l)
	return err
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}
