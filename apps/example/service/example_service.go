package service

import (
	"json-server-kit/apps/example/repository"
)

type ExampleService struct {
	repository *repository.ExampleRepository
}

func NewExampleService() *ExampleService {
	return &ExampleService{
		//repository,
	}
}

func (service *ExampleService) HelloWord() string {
	return "Hello Word!"
}
