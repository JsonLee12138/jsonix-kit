package dto

//go:generate easyjson -all ./dict.go

type DictTypeQuery struct {
	ID     *uint  `json:"id" form:"id"`
	Code   string `json:"code" form:"code"`
	Name   string `json:"name" form:"name"`
	Status *bool  `json:"status" form:"status"`
}
