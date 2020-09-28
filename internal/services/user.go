package services

import (
	"demo/internal/models/mysql/entity"
	"demo/internal/models/mysql/model"
)

func NewUser(u model.UserI) *User {
	return &User{
		UserModel: u,
	}
}

type UserI interface {
	Create(string, string, string) (entity.User, error)
	Get(string) (entity.User, error)
	Exist(string) (bool, error)
}

type User struct {
	UserModel model.UserI
}

func (u *User) Create(name, password, email string) (user entity.User, err error) {
	user, err = u.UserModel.Create(name, password, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) Get(name string) (user entity.User, err error) {
	user, err = u.UserModel.Get(name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) Exist(name string) (bool, error) {
	return u.UserModel.Exist(name)
}
