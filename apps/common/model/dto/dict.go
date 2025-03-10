package dto

type DictTypeQuery struct {
	ID     *uint  `json:"id" form:"id"`
	Code   string `json:"code" form:"code"`
	Name   string `json:"name" form:"name"`
	Status *bool  `json:"status" form:"status"`
}

type CreateDictTypeDTO struct {
	Code   string `json:"code" form:"code" validate:"required"`
	Name   string `json:"name" form:"name" validate:"required"`
	Status *bool  `json:"status" form:"status"`
	Desc   string `json:"desc" form:"desc"`
}
