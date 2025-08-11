package service

import (
	"CRUD_SERVER/repository"
	"CRUD_SERVER/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(userName string, newAge int64) error {
	return u.userRepository.Update(userName, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Name)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
