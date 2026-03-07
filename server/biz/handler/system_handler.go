package handler

import (
	"context"

	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type SystemHandler struct {
	confData *conf.Data
	svc      *service.Service
}

func NewSystemHandler(confData *conf.Data, svc *service.Service) api.SystemHandlerHTTPServer {
	return &SystemHandler{confData: confData, svc: svc}
}

func (h *SystemHandler) Initialize(ctx context.Context, req *api.InitializeRequest) (*api.InitializeResponse, error) {
	if req.GetUsername() != h.confData.SystemConfig.Username || req.GetPassword() != h.confData.SystemConfig.Password {
		return &api.InitializeResponse{Code: 1, Message: "认证失败"}, nil
	}
	adminUsername := req.GetAdminUsername()
	adminPassword := req.GetAdminPassword()
	if err := h.svc.SystemService.Initialize(ctx, adminUsername, adminPassword); err != nil {
		return &api.InitializeResponse{Code: 1, Message: err.Error()}, nil
	}
	return &api.InitializeResponse{Code: 0, Message: "success"}, nil
}
