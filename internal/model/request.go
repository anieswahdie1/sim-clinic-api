package model

import "github.com/asaskevich/govalidator"

type RequestPagination struct {
	Page   string `query:"page" valid:"required,alphanum"`
	Limit  string `query:"limit" valid:"required,alphanum"`
	Search string `query:"search" `
}

func (r *RequestPagination) Validate() error {
	_, err := govalidator.ValidateStruct(r)
	return err
}
