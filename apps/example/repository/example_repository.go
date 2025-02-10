package repository

import (
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{
		db,
	}
}

// func (r *ExampleRepository) FindAll() ([]Example, error) {
//     var examples []Example
//     result := r.db.Find(&examples)
//     return examples, result.Error
// }
