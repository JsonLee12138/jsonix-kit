package entity

import (
	"github.com/JsonLee12138/json-server/pkg/core"
	"gorm.io/gorm"
)

//go:generate easyjson -all ./user.go

// @AutoMigrate
type User struct {
	core.BaseEntityWithUuid
	Username string `json:"username" form:"username" gorm:"uniqueIndex;type:varchar(32)"` // 用户名
	Nickname string `json:"nickname" form:"nickname"`                                     // 昵称
	Phone    string `json:"phone" form:"phone" gorm:"index" validate:"phone"`             // 手机号
	Email    string `json:"email" form:"email" validate:"omitempty,email"`                // 邮箱
	Password string `json:"password,intern" form:"-" swaggerignore:"true"`                // 密码
	Enable   bool   `json:"enable" form:"enable" gorm:"index;default:1;type:tinyint(1)"`  // 是否启用
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.BaseEntityWithUuid.BeforeCreate(tx)
}
