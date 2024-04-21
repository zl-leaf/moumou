package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moumou/server/biz/handler"
)

func main() {
	router := gin.Default()
	router.Use(handler.CorsMW())
	handler.Register(router)
	err := router.Run(":8888")
	if err != nil {
		panic(err)
	}
}
