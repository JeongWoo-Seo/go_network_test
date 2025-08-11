package service

import (
	"CRUD_SERVER/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository

	User *User
}

func NewService(rep *repository.Repository) *Service {

	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
		serviceInstance.User = NewUserService(rep.User)
	})

	return serviceInstance
}
