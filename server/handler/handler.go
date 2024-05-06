package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/handler/vo"
	"github.com/moumou/server/util"
)

type Handler struct {
	svc *service.Service
}

func InitHandler() (*Handler, error) {
	var svc, err = service.InitService()
	if err != nil {
		return nil, err
	}
	return &Handler{
		svc: svc,
	}, nil
}

func (handler *Handler) Ping(ctx context.Context, req *vo.PingRequest) (resp *vo.PingResponse, err error) {
	resp = &vo.PingResponse{}
	return
}

func (handler *Handler) Hello(ctx context.Context, req *vo.HelloRequest) (resp *vo.HelloResponse, err error) {
	resp = &vo.HelloResponse{
		Data: &vo.HelloResponseData{
			Key: "1234567890abcdef",
			IV:  "abcdefghijklmnop",
		},
	}
	return
}

func (handler *Handler) Login(ctx context.Context, req *vo.LoginRequest) (resp *vo.LoginResponse, err error) {
	token, user, err := handler.svc.UserService.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	resp = &vo.LoginResponse{
		Data: &vo.LoginResponseData{
			User: &vo.UserData{
				UserID: util.Unit2String(user.ID),
			},
			Token: token,
		},
	}
	return
}

func (handler *Handler) Logout(ctx context.Context, req *vo.LogoutRequest) (resp *vo.LogoutResponse, err error) {
	err = handler.svc.UserService.Logout(req.Token)
	if err != nil {
		return
	}
	resp = &vo.LogoutResponse{}
	return
}

func (handler *Handler) Self(ctx context.Context, req *vo.SelfRequest) (resp *vo.SelfResponse, err error) {
	user, err := handler.svc.UserService.Self(req.Token)
	if err != nil {
		return nil, err
	}

	resp = &vo.SelfResponse{
		Data: &vo.UserData{
			UserID: util.Unit2String(user.ID),
		},
	}
	return
}

func (handler *Handler) RouterTree(ctx context.Context, request *vo.RouterTreeRequest) (resp *vo.RouterTreeResponse, err error) {
	routerList, err := handler.svc.PageService.GetRouterList()
	if err != nil {
		return nil, err
	}
	resp = &vo.RouterTreeResponse{
		Data: vo.ConvRouterList2RouterTreeData(routerList),
	}
	return
}
