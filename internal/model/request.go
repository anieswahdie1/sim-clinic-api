package model

import "github.com/asaskevich/govalidator"

type RequestPagination struct {
	Page   string `param:"page" json:"page" valid:"required,alphanum"`
	Limit  string `param:"limit" json:"limit" valid:"required,alphanum"`
	Search string `param:"search" json:"search"`
}

func (r *RequestPagination) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}
