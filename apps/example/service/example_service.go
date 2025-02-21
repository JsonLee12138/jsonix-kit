package service

import (
	"json-server-kit/apps/example/repository"
)

type ExampleService struct {
    repository *repository.ExampleRepository
}

func NewExampleService(repository *repository.ExampleRepository) *ExampleService {
	return &ExampleService{
	    repository,
	}
}

func (service *ExampleService) HelloWorld() string {
	return "Hello World!"
}
