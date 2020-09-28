// Package apis provides ...
package apis

import (
	"demo/internal/models/mysql/entity"
	"fmt"
	"net/http"

	"demo/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserI services.UserI
}

func (u *User) CreateUser(c *gin.Context) {
	var user entity.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, "err")
		return
	}
	fmt.Println(user)
	if user.UserName == "" {
		c.JSON(400, "bad request")
		return
	}

	ok, err := u.UserI.Exist(user.UserName)
	fmt.Println(ok, err)
	if err != nil {

	}
	if ok {
		c.JSON(500, "error")
		return
	}

	userS, err := u.UserI.Create(user.UserName, user.Password, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
	}
	c.JSON(0, userS)
}

func (u *User) GetUser(c *gin.Context) {
	name, _ := c.GetQuery("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println(name)
	user, err := u.UserI.Get(name)
	fmt.Println(err)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, "not found")
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}

	c.JSON(0, user)
	return
}
