package handler

import (
	"context"
	"errors"
)

var errorResponse = &BaseResponse{
	Code:    -1,
	Message: "error",
}

func Ping(ctx context.Context, req *PingRequest) (resp *PingResponse, err error) {
	resp = &PingResponse{}
	return
}

func Hello(ctx context.Context, req *HelloRequest) (resp *HelloResponse, err error) {
	resp = &HelloResponse{
		Data: &HelloResponseData{
			Key: "1234567890abcdef",
			IV:  "abcdefghijklmnop",
		},
	}
	return
}

func Login(ctx context.Context, req *LoginRequest) (resp *LoginResponse, err error) {
	if req.Username != "demo" || req.Password != "uuAnx9IgNja9Eyw0kojAnA==" {
		err = errors.New("账号或密码错误")
		return
	}

	resp = &LoginResponse{
		Data: &LoginResponseData{
			User: &UserData{
				UserID: "111111",
			},
			Token: "token",
		},
	}
	return
}

func Logout(ctx context.Context, req *LogoutRequest) (resp *LogoutResponse, err error) {
	if req.Token != "token" {
		err = errors.New("token不存在")
		return
	}
	resp = &LogoutResponse{}
	return
}

func Self(ctx context.Context, req *SelfRequest) (resp *SelfResponse, err error) {
	if req.Token != "token" {
		return nil, errors.New("找不到用户")
	}
	resp = &SelfResponse{
		Data: &UserData{
			UserID: "111111",
		},
	}
	return
}

func Menu(ctx context.Context, request *MenuRequest) (resp *MenuResponse, err error) {
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
