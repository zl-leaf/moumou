package mw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Method", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "OK")
			return
		}
		c.Next()
	}
}
