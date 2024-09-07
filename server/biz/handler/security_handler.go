package handler

import (
	"context"

	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type SecurityHandler struct {
	svc *service.Service
}

func NewSecurityHandler(svc *service.Service) api.SecurityHandlerHTTPServer {
	return &SecurityHandler{svc}
}

func (h *SecurityHandler) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	token, user, err := h.svc.UserService.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &api.LoginResponse{
		Data: &api.LoginResponseData{
			Token: token,
			User:  ConvUser2VO(user),
		},
	}, nil
}

func (h *SecurityHandler) Logout(ctx context.Context, request *api.LogoutRequest) (*api.LogoutResponse, error) {
	err := h.svc.UserService.Logout(ctx, "token")
	if err != nil {
		return nil, err
	}
	return &api.LogoutResponse{}, nil
}

func (h *SecurityHandler) Self(ctx context.Context, request *api.SelfRequest) (*api.SelfResponse, error) {
	user, err := h.svc.UserService.Self(ctx)
	if err != nil {
		return nil, err
	}
	return &api.SelfResponse{
		Data: ConvUser2VO(user),
	}, nil
}
