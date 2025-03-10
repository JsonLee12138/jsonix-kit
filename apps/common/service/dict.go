package service

import (
	"json-server-kit/apps/common/entity"
	"json-server-kit/apps/common/model/dto"
	"json-server-kit/apps/common/repository"
)

type DictService struct {
	dictRepository *repository.DictRepository
}

func NewDictService(dictRepository *repository.DictRepository) *DictService {
	return &DictService{
		dictRepository,
	}
}

func (s *DictService) GetDictTypes(params dto.DictTypeQuery) ([]entity.DictType, error) {
	return s.dictRepository.GetDictTypes(params)
}

func (s *DictService) CreateDictType(body dto.CreateDictTypeDTO) (entity.DictType, error) {
	return s.dictRepository.CreateDictType(body)
}

func (s *DictService) GetDictItems(code string) ([]entity.DictItem, error) {
	return s.dictRepository.GetDictItemsByType(code)
}
