package handler

import (
	"context"

	"github.com/moumou/server/biz/conv"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type SecurityHandler struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewSecurityHandler(svc *service.Service, converter conv.IConverter) api.SecurityHandlerHTTPServer {
	return &SecurityHandler{svc: svc, converter: converter}
}

func (h *SecurityHandler) Captcha(ctx context.Context, request *api.CaptchaRequest) (*api.CaptchaResponse, error) {
	randomId, image, err := h.svc.UserService.RandomCaptcha()
	if err != nil {
		return nil, err
	}

	return &api.CaptchaResponse{
		Data: &api.CaptchaResponseData{
			RandomId: randomId,
			Image:    image,
		},
	}, nil
}

func (h *SecurityHandler) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	err := h.svc.UserService.VerifyCaptcha(ctx, request.GetCaptchaId(), request.GetCaptcha())
	if err != nil {
		return nil, err
	}

	token, user, err := h.svc.UserService.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &api.LoginResponse{
		Data: &api.LoginResponseData{
			Token: token,
			User:  h.converter.ConvertUserToVO(user),
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
		Data: h.converter.ConvertUserToVO(user),
	}, nil
}
