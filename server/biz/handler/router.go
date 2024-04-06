package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		ctx := context.TODO()
		req := &PingRequest{}
		resp, err := Ping(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/hello", func(c *gin.Context) {
		ctx := context.TODO()
		req := &HelloRequest{}

		resp, err := Hello(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/login", func(c *gin.Context) {
		ctx := context.TODO()
		req := &LoginRequest{}
		err := c.BindJSON(req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := Login(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/logout", func(c *gin.Context) {
		ctx := context.TODO()
		req := &LogoutRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := Logout(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/self", func(c *gin.Context) {
		ctx := context.TODO()
		req := &SelfRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := Self(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/menu", func(c *gin.Context) {
		ctx := context.TODO()
		req := &MenuRequest{}
		/*	err := c.BindJSON(req)
			if err != nil {
				c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: "参数错误"})
				return
			}*/

		resp, err := Menu(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
