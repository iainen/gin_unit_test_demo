package main

import (
	"demo/internal/models/mysql"
	"demo/internal/router"
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
