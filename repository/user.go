package repository

import (
	"CRUD_SERVER/types"
	"CRUD_SERVER/types/errors"
)

type UserRepository struct {
	UserMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		UserMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.UserMap = append(u.UserMap, newUser)
	return nil
}

func (u *UserRepository) Update(userName string, newAge int64) error {
	isExisted := false

	for i := len(u.UserMap) - 1; i >= 0; i-- {
		if u.UserMap[i].Name == userName {
			u.UserMap[i].Age = newAge
			isExisted = true
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	}

	return nil
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false

	for i := len(u.UserMap) - 1; i >= 0; i-- {
		if u.UserMap[i].Name == userName {
			u.UserMap = append(u.UserMap[:i], u.UserMap[i+1:]...)
			isExisted = true
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	}

	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.UserMap
}
