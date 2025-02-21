package entity

import (
	"github.com/JsonLee12138/jsonix/pkg/core"
	"gorm.io/gorm"
)

//go:generate easyjson -all ./

// @AutoMigrate
type ExampleEntity struct {
	core.BaseEntityWithUuid
}

func (e *ExampleEntity) BeforeCreate(tx *gorm.DB) error {
	return e.BaseEntityWithUuid.BeforeCreate(tx)
}
