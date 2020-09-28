/**
* @author : yi.zhang
* @description : model 描述
* @date   : 2020-08-17 18:27
 */

package model

import (
	"demo/internal/models/mysql/entity"

	"github.com/jinzhu/gorm"
)

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

type UserI interface {
	Create(string, string, string) (entity.User, error)
	Get(string) (entity.User, error)
	Exist(string) (bool, error)
}

type User struct {
	DB *gorm.DB
}

func (u *User) Create(name string, password string, email string) (user entity.User, err error) {
	user.UserName = name
	user.Password = password
	user.Email = email
	err = u.DB.Create(&user).Error
	return user, nil
}

func (u *User) Get(name string) (user entity.User, err error) {
	err = u.DB.Where("user_name =?", name).Find(&user).Error
	return user, err
}

func (u *User) Exist(name string) (bool, error) {
	var count int
	err := u.DB.Model(&entity.User{}).Where("user_name=?", name).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
