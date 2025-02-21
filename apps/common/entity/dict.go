package entity

import "github.com/JsonLee12138/jsonix/pkg/core"

//go:generate easyjson -all ./dict.go

// @AutoMigrate
type DictType struct {
	core.BaseEntityWithID
	Code   string     `json:"code" form:"code" gorm:"uniqueIndex;type:varchar(32)"`        // 字典编码
	Name   string     `json:"name" form:"name" gorm:"type:varchar(128)"`                   // 字典名称
	Status *bool      `json:"status" form:"status" gorm:"index;default:1;type:tinyint(1)"` // 状态
	Desc   string     `json:"desc" form:"desc" gorm:"type:varchar(255)"`                   // 描述
	Items  []DictItem `json:"items" form:"items"`
}

// @AutoMigrate
type DictItem struct {
	core.BaseEntityWithID
	Key        string `json:"key" form:"key" gorm:"uniqueIndex;type:varchar(50)"`                                               // 字典项键
	Value      string `json:"value" form:"value" gorm:"type:varchar(100)"`                                                      // 字典项值
	Name       string `json:"name" form:"name" gorm:"type:varchar(100)"`                                                        // 字典项名称
	Desc       string `json:"desc" form:"desc" gorm:"type:varchar(255)"`                                                        // 字典项描述
	Sort       int    `json:"sort" form:"sort" gorm:"default:0;type:int(11)"`                                                   // 排序
	Status     *bool  `json:"status" form:"status" gorm:"index;default:1;type:tinyint(1)"`                                      // 状态(1: 正常, 0: 禁用)
	ClassList  string `json:"classList" form:"classList" gorm:"type:varchar(255)"`                                              // 样式列表
	Color      string `json:"color" form:"color" gorm:"type:varchar(20)"`                                                       // 颜色
	IsDefault  *bool  `json:"isDefault" form:"isDefault" gorm:"default:0;type:tinyint(1)"`                                      // 是否默认
	DictTypeId uint   `json:"dictTypeId" form:"dictTypeId" gorm:"index;column:dic_type_id;foreignKey:DictTypeId;references:ID"` // 字典类型ID
}
