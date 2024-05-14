package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/moumou/server/handler/vo"
	"net/http"
)

func Register(engine *gin.Engine) error {
	var handler, err = InitHandler()
	if err != nil {
		return err
	}
	engine.GET("/ping", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.PingRequest{}
		resp, err := handler.Ping(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/hello", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.HelloRequest{}

		resp, err := handler.Hello(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/login", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.LoginRequest{}
		err := c.BindJSON(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := handler.Login(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/logout", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.LogoutRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := handler.Logout(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/self", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.SelfRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := handler.Self(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/router_tree", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.RouterTreeRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}

		resp, err := handler.RouterTree(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/page", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.GetPageRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}
		err = c.BindJSON(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}
		resp, err := handler.GetPage(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	engine.POST("/page_data_list", func(c *gin.Context) {
		ctx := context.TODO()
		req := &vo.GetPageDataListRequest{}
		err := c.BindHeader(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}
		err = c.BindJSON(req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: "参数错误"})
			return
		}
		resp, err := handler.GetPageDataList(ctx, req)
		if err != nil {
			c.JSON(http.StatusOK, &vo.BaseResponse{Code: -1, Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	return nil
}
