package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moumou/server/framework/config"
	"github.com/moumou/server/handler"
	"github.com/moumou/server/handler/mw"
)

type serverConfig struct {
	Port string `yaml:"port"`
}

func main() {
	router := gin.Default()
	router.Use(mw.CorsMW())

	var err error
	err = config.Init("config")
	if err != nil {
		panic(err)
	}
	var cnf serverConfig
	err = config.GetConfig("server", &cnf)
	if err != nil {
		panic(err)
	}

	err = handler.Register(router)
	if err != nil {
		panic(err)
	}

	err = router.Run(":" + cnf.Port)
	if err != nil {
		panic(err)
	}
}
