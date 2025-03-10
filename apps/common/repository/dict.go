package repository

import (
	"json-server-kit/apps/common/entity"
	"json-server-kit/apps/common/model/dto"

	"gorm.io/gorm"
)

type DictRepository struct {
	db *gorm.DB
}

func NewDictRepository(db *gorm.DB) *DictRepository {
	return &DictRepository{db: db}
}

func (r *DictRepository) GetDictTypes(params dto.DictTypeQuery) ([]entity.DictType, error) {
	var dictTypes []entity.DictType
	db := r.db.Model(&entity.DictType{})
	if params.ID != nil {
		db = db.Where("id = ?", params.ID)
	}
	if params.Code != "" {
		db = db.Where("code = ?", params.Code)
	}
	if params.Name != "" {
		db = db.Where("name = ?", params.Name)
	}
	if params.Status != nil {
		db = db.Where("status = ?", params.Status)
	}
	if err := db.Find(&dictTypes).Error; err != nil {
		return nil, err
	}
	return dictTypes, nil
}

func (r *DictRepository) CreateDictType(body dto.CreateDictTypeDTO) (entity.DictType, error) {
	dictType := entity.DictType{
		Code:   body.Code,
		Name:   body.Name,
		Status: body.Status,
		Desc:   body.Desc,
	}
	if err := r.db.Create(&dictType).Error; err != nil {
		return entity.DictType{}, err
	}
	return dictType, nil
}

func (r *DictRepository) GetDictTypeById(id uint) (entity.DictType, error) {
	var dictType entity.DictType
	if err := r.db.Where("id = ?", id).First(&dictType).Preload("Items").Error; err != nil {
		return entity.DictType{}, err
	}
	return dictType, nil
}

func (r *DictRepository) GetDictItemsByType(code string) ([]entity.DictItem, error) {
	var items []entity.DictItem
	if err := r.db.Model(&entity.DictItem{}).Select("dict_items.*").Joins("inner join dict_types dt on dict_items.dict_type_id = dt.id").Where("dt.code = ?", code).Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
