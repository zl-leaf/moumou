package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/proto_gen"
)

type SecurityHandler struct {
	svc *service.Service
}

func NewSecurityHandler(svc *service.Service) api.SecurityHandlerHTTPServer {
	return &SecurityHandler{svc}
}

func (h *SecurityHandler) GetPublicKey(ctx context.Context, request *api.GetPublicKeyRequest) (*api.GetPublicKeyResponse, error) {
	return &api.GetPublicKeyResponse{
		Data: &api.GetPublicKeyResponseData{
			Key: "1234567890abcdef",
			Iv:  "abcdefghijklmnop",
		},
	}, nil
}

func (h *SecurityHandler) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	token, user, err := h.svc.SysUserService.Login(request.GetUsername(), request.GetPassword())
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
	err := h.svc.SysUserService.Logout("token")
	if err != nil {
		return nil, err
	}
	return &api.LogoutResponse{}, nil
}

func (h *SecurityHandler) Self(ctx context.Context, request *api.SelfRequest) (*api.SelfResponse, error) {
	user, err := h.svc.SysUserService.Self("token")
	if err != nil {
		return nil, err
	}
	return &api.SelfResponse{
		Data: ConvUser2VO(user),
	}, nil
}

func (h *SecurityHandler) GetSecurityRouterTree(ctx context.Context, request *api.GetSecurityRouterTreeRequest) (*api.GetSecurityRouterTreeResponse, error) {
	routerList, err := h.svc.PageService.GetRouterList()
	if err != nil {
		return nil, err
	}

	return &api.GetSecurityRouterTreeResponse{
		Data: ConvRouterList2SecurityRouterTreeRespData(routerList),
	}, nil
}
