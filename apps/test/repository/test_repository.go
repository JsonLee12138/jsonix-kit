package repository

import (
    "gorm.io/gorm"
)

type TestRepository struct {
    db *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
    return &TestRepository{
        db,
    }
}
