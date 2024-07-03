package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/handler/vo"
)

type SysUserManageHandler struct {
	svc *service.Service
}

func (h *SysUserManageHandler) GetSysUserList(ctx context.Context, request *vo.GetSysUserListRequest) (resp *vo.GetSysUserListResponse, err error) {
	sysUserList, total, err := h.svc.SysUserService.GetSysUserList(ctx, &request.Filter.SysUserFilter, request.CurrentPage, request.PageSize)
	if err != nil {
		return nil, err
	}
	resp = &vo.GetSysUserListResponse{
		Data: vo.ConvSysUserList2ResponseData(sysUserList, total),
	}
	return
}

func (h *SysUserManageHandler) getSysUserInfo(ctx context.Context, request *vo.GetSysUserInfoRequest) (resp *vo.GetSysUserInfoResponse, err error) {
	sysUserInfo, err := h.svc.SysUserService.GetSysUserInfo(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	resp = &vo.GetSysUserInfoResponse{
		Data: vo.ConvSysUser2VO(sysUserInfo),
	}
	return
}
