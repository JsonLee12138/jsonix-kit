package entity

import "github.com/JsonLee12138/json-server/pkg/core"

//go:generate easyjson -all ./user_entity.go

type User struct {
	core.BaseEntityWithUuid
	Username string `json:"username" form:"username" gorm:"uniqueIndex;type:varchar(32)"`
	Nickname string `json:"nickname" form:"nickname"`
	Phone    string `json:"phone" form:"phone" gorm:"index" validate:"phone"`
	Email    string `json:"email" form:"email" validate:"omitempty,email"`
	// RoleIds  []uint `json:"-" gorm:"-"`
	// Roles    []Role    `json:"roles" gorm:"many2many:user_roles;joinForeignKey:UserId;joinReferences:RoleId"`
	// DeptId uint `json:"deptId" gorm:"index"`
	// //Dept          *Dept     `json:"dept" gorm:"foreignKey:DeptId;references:ID"`
	// JobPositionId *uint  `json:"jobPositionId" gorm:"index"`
	Password string `json:"-" form:"-"`
	Enable   bool   `json:"enable" form:"enable" gorm:"index;default:1;type:tinyint(1)"`
}

func (u *User) BeforeCreate() error {
	return u.BaseEntityWithUuid.BeforeCreate()
}
