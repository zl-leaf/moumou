package handler

import (
	"context"

	"github.com/moumou/server/biz/conv"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/biz/util/ctxutil"
	api "github.com/moumou/server/gen/proto"
)

type PermissionHandler struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewPermissionHandler(
	svc *service.Service,
	converter conv.IConverter,
) api.PermissionHandlerHTTPServer {
	return &PermissionHandler{svc: svc, converter: converter}
}

func (p PermissionHandler) DeletePermission(ctx context.Context, request *api.DeletePermissionRequest) (*api.DeletePermissionResponse, error) {
	err := p.svc.Dao.PermissionDao(ctx).Delete(request.Ids)
	if err != nil {
		return nil, err
	}
	return &api.DeletePermissionResponse{}, nil
}

func (p PermissionHandler) GetPermissionInfo(ctx context.Context, request *api.GetPermissionInfoRequest) (*api.GetPermissionInfoResponse, error) {
	permission, err := p.svc.Dao.PermissionDao(ctx).GetByID(request.GetId())
	if err != nil {
		return nil, err
	}

	return &api.GetPermissionInfoResponse{
		Data: p.converter.ConvertPermissionToVO(permission),
	}, nil
}

func (p PermissionHandler) GetPermissionTree(ctx context.Context, request *api.GetPermissionTreeRequest) (*api.GetPermissionTreeResponse, error) {
	topLevelPermissions, err := p.svc.RoleService.GetTopLevelPermissions(ctx)
	if err != nil {
		return nil, err
	}

	return &api.GetPermissionTreeResponse{
		Data: &api.GetPermissionTreeResponseData{
			List: p.converter.ConvertPermissionTreeNodeListToVO(topLevelPermissions),
		},
	}, nil
}

func (p PermissionHandler) CreatePermission(ctx context.Context, request *api.CreatePermissionRequest) (*api.CreatePermissionResponse, error) {
	permission := new(model.Permission)
	p.converter.ConvertCreatePermissionRequestDataToBO(request.GetPermission(), permission)
	err := p.svc.Dao.PermissionDao(ctx).Create(permission)
	if err != nil {
		return nil, err
	}

	return &api.CreatePermissionResponse{
		Data: &api.CreatePermissionResponseData{Id: permission.Id},
	}, nil
}

func (p PermissionHandler) UpdatePermission(ctx context.Context, request *api.UpdatePermissionRequest) (*api.UpdatePermissionResponse, error) {
	permission, err := p.svc.Dao.PermissionDao(ctx).GetByID(request.GetPermission().GetId())
	if err != nil {
		return nil, err
	}
	p.converter.ConvertUpdatePermissionRequestDataToBO(request.GetPermission(), permission)
	err = p.svc.Dao.PermissionDao(ctx).Save(permission)
	if err != nil {
		return nil, err
	}
	return &api.UpdatePermissionResponse{}, nil
}

func (p PermissionHandler) GetPermissionList(ctx context.Context, request *api.GetPermissionListRequest) (*api.GetPermissionListResponse, error) {
	permissionList, total, err := p.svc.Dao.PermissionDao(ctx).OrderBySort().Find()
	if err != nil {
		return nil, err
	}
	return &api.GetPermissionListResponse{
		Data: &api.GetPermissionListResponseData{
			Total: total,
			List:  p.converter.ConvertPermissionListToVO(permissionList),
		},
	}, nil
}

func (p PermissionHandler) GetUserPermission(ctx context.Context, request *api.GetUserPermissionRequest) (*api.GetUserPermissionResponse, error) {
	selfUserId := ctxutil.GetCtxUserID(ctx)
	permissions, err := p.svc.RoleService.GetPermissionCodesByUid(ctx, selfUserId)
	if err != nil {
		return nil, err
	}

	permissionCodes := make([]string, len(permissions))
	for i, permission := range permissions {
		permissionCodes[i] = permission.Code
	}

	return &api.GetUserPermissionResponse{
		Data: &api.GetUserPermissionResponseData{
			Permissions: permissionCodes,
		},
	}, nil
}
