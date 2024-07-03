package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/handler/vo"
)

type SysRouterManageHandler struct {
	svc *service.Service
}

func (h *SysRouterManageHandler) GetSysRouterList(ctx context.Context, req *vo.GetSysRouterListRequest) (resp *vo.GetSysRouterListResponse, err error) {
	routerList, err := h.svc.SysRouterService.GetSysRouterList(ctx)
	if err != nil {
		return nil, err
	}
	resp = &vo.GetSysRouterListResponse{
		Data: vo.ConvSysRouterList2ResponseData(routerList),
	}
	return
}

func (h *SysRouterManageHandler) GetSysRouterInfo(ctx context.Context, req *vo.GetSysRouterInfoRequest) (resp *vo.GetSysRouterInfoResponse, err error) {
	routerInfo, err := h.svc.SysRouterService.GetSysRouterInfo(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	resp = &vo.GetSysRouterInfoResponse{
		Data: vo.ConvSysRouter2VO(routerInfo),
	}
	return
}
