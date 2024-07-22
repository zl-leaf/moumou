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

	return &api.GetRouterListResponse{
		Data: ConvRouterList2ResponseData(routerList),
	}, nil
}

func (h *RouterHandler) CreateRouter(ctx context.Context, request *api.CreateRouterRequest) (*api.CreateRouterResponse, error) {
	routerID, err := h.svc.SysRouterService.CreateRouter(ctx, ConvCreateRequestData2FormData(request.Router))
	if err != nil {
		return nil, err
	}
	return &api.CreateRouterResponse{
		Data: &api.CreateRouterResponseData{
			Id: routerID,
		},
	}, nil
}

func (h *RouterHandler) UpdateRouter(ctx context.Context, request *api.UpdateRouterRequest) (*api.UpdateRouterResponse, error) {
	err := h.svc.SysRouterService.UpdateRouter(ctx, ConvUpdateRequestData2FormData(request.Router))
	if err != nil {
		return nil, err
	}
	return &api.UpdateRouterResponse{}, nil
}

func (h *RouterHandler) GetRouterInfo(ctx context.Context, request *api.GetRouterInfoRequest) (*api.GetRouterInfoResponse, error) {
	router, err := h.svc.SysRouterService.GetRouterInfo(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &api.GetRouterInfoResponse{
		Data: ConvRouter2VO(router),
	}, nil
}

func (h *RouterHandler) DeleteRouter(ctx context.Context, request *api.DeleteRouterRequest) (*api.DeleteRouterResponse, error) {
	err := h.svc.SysRouterService.DeleteRouter(ctx, request.GetIds())
	if err != nil {
		return nil, err
	}
	return &api.DeleteRouterResponse{}, nil
}
