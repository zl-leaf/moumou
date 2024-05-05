package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(engine *gin.Engine) error {
	var handler, err = InitHandler()
	if err != nil {
		return err
	}
	engine.GET("/ping", func(c *gin.Context) {
		ctx := context.TODO()
		req := &PingRequest{}
		resp, err := handler.Ping(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/hello", func(c *gin.Context) {
		ctx := context.TODO()
		req := &HelloRequest{}

		resp, err := handler.Hello(ctx, req)
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

		resp, err := handler.Login(ctx, req)
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

		resp, err := handler.Logout(ctx, req)
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

		resp, err := handler.Self(ctx, req)
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

		resp, err := handler.Menu(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	return nil
}
