package entity

import (
	"github.com/JsonLee12138/json-server/pkg/core"
)

type TestEntity struct {
	core.BaseEntityWithUuid
}

func (e *TestEntity) BeforeCreate() error {
	return e.BaseEntityWithUuid.BeforeCreate()
}
