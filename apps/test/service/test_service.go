package service

import (
	"json-server-kit/apps/test/repository"
)

type TestService struct {
    repository *repository.TestRepository
}

func NewTestService(repository *repository.TestRepository) *TestService {
	return &TestService{
	    repository,
	}
}

func (service *TestService) HelloWord() string {
	return "Hello Word!"
}