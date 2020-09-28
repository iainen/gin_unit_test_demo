/**
* @author : yi.zhang
* @description : gin_demo 描述
* @date   : 2020-05-07 11:36
 */

package main

import (
	"demo/internal/models/mysql"
	"demo/internal/router"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	mysql.Init()
	db := mysql.GetDB()

	ro := router.InitRouter(db)
	router.Init(ro)
}

func main() {
	fmt.Println("main")
	gin.SetMode("debug")

	r := router.Get()

	server := &http.Server{
		Addr:    ":4442",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		os.Exit(1)
	}
}
