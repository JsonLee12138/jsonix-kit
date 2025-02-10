package service

import (
	"json-server-kit/apps/auth/repository"
)

type UsernameService struct {
	repository *repository.UserRepository
}

func NewExampleService(repo *repository.UserRepository) *UsernameService {
	return &UsernameService{
		repository: repo,
	}
}

func (service *UsernameService) HelloWord() string {
	return "Hello Word!"
}
