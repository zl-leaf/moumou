package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moumou/server/handler"
)

func main() {
	router := gin.Default()
	router.Use(handler.CorsMW())

	var err error
	err = handler.Register(router)
	if err != nil {
		panic(err)
	}
	err = router.Run(":8888")
	if err != nil {
		panic(err)
	}
}
