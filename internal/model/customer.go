package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"time"
)

func init() {
	govalidator.TagMap["gender"] = govalidator.Validator(func(str string) bool {
		return str == "L" || str == "P"
	})
}

type Customer struct {
	Id                 string         `json:"id" gorm:"primary_key;unique"`
	CodeRegister       string         `json:"codeRegister" gorm:"code_register"`
	CustomerName       string         `json:"customerName" gorm:"customer_name"`
	PhoneNumber        string         `json:"phoneNumber" gorm:"phone_number"`
	CustomerAddress    string         `json:"customerAddress" gorm:"customer_address"`
	Gender             string         `json:"gender" gorm:"gender"`
	InformedConsent    string         `json:"informedConsent" gorm:"informed_consent"`
	SourceTerapistInfo string         `json:"sourceTerapistInfo" gorm:"source_terapist_info"`
	City               string         `json:"city" gorm:"city"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type AddCustomerRequest struct {
	CustomerName       string `json:"customerName" valid:"required,length(3|100)"`
	PhoneNumber        string `json:"phoneNumber" valid:"required,length(10|13)"`
	CustomerAddress    string `json:"customerAddress"`
	Gender             string `json:"gender" valid:"required,gender"`
	InformedConsent    string `json:"informedConsent"`
	SourceTerapistInfo string `json:"sourceTerapistInfo"`
	City               string `json:"city"`
}

func (r AddCustomerRequest) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}
