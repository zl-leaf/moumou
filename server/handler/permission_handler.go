package handler

import (
	"context"

	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type PermissionHandler struct {
	svc *service.Service
}

func NewPermissionHandler(svc *service.Service) api.PermissionHandlerHTTPServer {
	return &PermissionHandler{svc: svc}
}

func (p PermissionHandler) GetPermissionList(ctx context.Context, request *api.GetPermissionListRequest) (*api.GetPermissionListResponse, error) {
	permissionList, total, err := p.svc.Dao.PermissionDao.OrderBySort().Find()
	if err != nil {
		return nil, err
	}
	return &api.GetPermissionListResponse{
		Data: &api.GetPermissionListResponseData{
			Total: total,
			List:  ConvPermissionList2VO(permissionList),
		},
	}, nil
}
