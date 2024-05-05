package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
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

func (handler *Handler) Ping(ctx context.Context, req *PingRequest) (resp *PingResponse, err error) {
	resp = &PingResponse{}
	return
}

func (handler *Handler) Hello(ctx context.Context, req *HelloRequest) (resp *HelloResponse, err error) {
	resp = &HelloResponse{
		Data: &HelloResponseData{
			Key: "1234567890abcdef",
			IV:  "abcdefghijklmnop",
		},
	}
	return
}

func (handler *Handler) Login(ctx context.Context, req *LoginRequest) (resp *LoginResponse, err error) {
	token, user, err := handler.svc.UserService.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	resp = &LoginResponse{
		Data: &LoginResponseData{
			User: &UserData{
				UserID: util.Unit2String(user.ID),
			},
			Token: token,
		},
	}
	return
}

func (handler *Handler) Logout(ctx context.Context, req *LogoutRequest) (resp *LogoutResponse, err error) {
	err = handler.svc.UserService.Logout(req.Token)
	if err != nil {
		return
	}
	resp = &LogoutResponse{}
	return
}

func (handler *Handler) Self(ctx context.Context, req *SelfRequest) (resp *SelfResponse, err error) {
	user, err := handler.svc.UserService.Self(req.Token)
	if err != nil {
		return nil, err
	}

	resp = &SelfResponse{
		Data: &UserData{
			UserID: util.Unit2String(user.ID),
		},
	}
	return
}

func (handler *Handler) Menu(ctx context.Context, request *MenuRequest) (resp *MenuResponse, err error) {
	resp = &MenuResponse{
		Data: &MenuResponseData{
			MenuItems: []MenuItem{
				{
					Name:     "home_welcome",
					Path:     "/welcome",
					Title:    "你好",
					IsMenu:   true,
					Children: []MenuItem{},
				},
				{
					Name:   "home_group",
					Path:   "welcome",
					Title:  "菜单",
					IsMenu: true,
					Children: []MenuItem{
						{
							Name:     "home_demo_list",
							Path:     "/demo_list",
							Title:    "列表",
							IsMenu:   true,
							Children: []MenuItem{},
						},
					},
				},
			},
		},
	}
	return
}
