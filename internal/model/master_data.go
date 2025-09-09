package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

type LayananTerapi struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Code      string         `json:"code" gorm:"uniqueIndex;not null" valid:"required,alphanum,length(3|20)"`
	Name      string         `json:"name" gorm:"not null" valid:"required,length(3|100)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RiwayatPenyakit struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null" valid:"required,alphanum,length(3|20)"`
	Name        string         `json:"name" gorm:"not null" valid:"required,length(3|100)"`
	Description string         `json:"description" gorm:"type:text" valid:"optional,length(0|500)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type TeknikTerapi struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null" valid:"required,alphanum,length(3|20)"`
	Name        string         `json:"name" gorm:"not null" valid:"required,length(3|100)"`
	Description string         `json:"description" gorm:"type:text" valid:"optional,length(0|500)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type LayananTerapiRequest struct {
	Code string `json:"code" valid:"required,alphanum,length(3|20)"`
	Name string `json:"name" valid:"required,length(3|100)"`
}

type RiwayatPenyakitRequest struct {
	Code        string `json:"code" valid:"required,alphanum,length(3|20)"`
	Name        string `json:"name" valid:"required,length(3|100)"`
	Description string `json:"description" valid:"optional,length(0|500)"`
}

type TeknikTerapiRequest struct {
	Code        string `json:"code" valid:"required,alphanum,length(3|20)"`
	Name        string `json:"name" valid:"required,length(3|100)"`
	Description string `json:"description" valid:"optional,length(0|500)"`
}

func (l *LayananTerapi) Validate() error {
	_, err := govalidator.ValidateStruct(l)
	return err
}

func (r *RiwayatPenyakit) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}

func (t *TeknikTerapi) Validate() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}

func (r *LayananTerapiRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}

func (r *RiwayatPenyakitRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}

func (r *TeknikTerapiRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}
