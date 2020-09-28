/**
* @author : yi.zhang
* @description : router 描述
* @date   : 2020-08-17 16:31
 */

package router

import (
	"net/http"

	"demo/internal/apis"
	"demo/internal/models/mysql/model"
	"demo/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var r = gin.New()

type Router struct {
	userApi *apis.User
}

func InitRouter(db *gorm.DB) *Router {
	user := &model.User{
		DB: db,
	}

	serviceUser := &services.User{
		UserModel: model.UserI(user),
	}

	apiUser := &apis.User{
		UserI: services.UserI(serviceUser),
	}

	var ro Router
	ro.userApi = apiUser

	return &ro
}

// Init ...
func Init(ro *Router) {
	r = gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "the incorrect apis route")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST(UriApiCreateUser, ro.userApi.CreateUser)
	r.GET(UriApiGetUser, ro.userApi.GetUser)
}

func Get() *gin.Engine {
	return r
}
