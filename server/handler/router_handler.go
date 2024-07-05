package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/proto_gen"
)

type RouterHandler struct {
	svc *service.Service
}

func NewRouterHandler(svc *service.Service) api.RouterHandlerHTTPServer {
	return &RouterHandler{svc: svc}
}

func (h *RouterHandler) GetRouterList(ctx context.Context, request *api.GetRouterListRequest) (*api.GetRouterListResponse, error) {
	routerList, err := h.svc.SysRouterService.GetRouterList(ctx)
	if err != nil {
		return nil, err
	}

	//token := ctx.(http.Context).Header().Get("x-token")
	return &api.GetRouterListResponse{
		Data: ConvRouterList2ResponseData(routerList),
	}, nil
}

func (h *RouterHandler) CreateRouter(context.Context, *api.CreateRouterRequest) (*api.CreateRouterResponse, error) {
	return &api.CreateRouterResponse{}, nil
}

func (h *RouterHandler) GetRouterInfo(ctx context.Context, request *api.GetRouterInfoRequest) (*api.GetRouterInfoResponse, error) {
	router, err := h.svc.SysRouterService.GetRouterInfo(ctx, int(request.GetId()))
	if err != nil {
		return nil, err
	}

	return &api.GetRouterInfoResponse{
		Data: ConvRouter2VO(router),
	}, nil
}
