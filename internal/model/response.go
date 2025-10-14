package model

type ResponsePagination struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Data  any   `json:"data"`
}
