package handler

import (
	"context"

	"github.com/moumou/server/biz/conv"

	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type PermissionHandler struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewPermissionHandler(svc *service.Service, converter conv.IConverter) api.PermissionHandlerHTTPServer {
	return &PermissionHandler{svc: svc, converter: converter}
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

func (p PermissionHandler) GetUserPermissionPath(ctx context.Context, request *api.GetUserPermissionPathRequest) (*api.GetUserPermissionPathResponse, error) {
	user, err := p.svc.UserService.Self(ctx)
	if err != nil {
		return nil, err
	}
	userRoleList, _, err := p.svc.Dao.UserRelRoleDao(ctx).WhereUserIDEq(user.Id).Find()
	roleIDs := make([]int64, len(userRoleList))
	for i, userRole := range userRoleList {
		roleIDs[i] = userRole.RoleID
	}
	rolePermissionList, _, err := p.svc.Dao.RolePermissionDao(ctx).WhereRoleIDIn(roleIDs).Find()
	if err != nil {
		return nil, err
	}
	permissionIDList := make([]int64, len(rolePermissionList))
	for i, rolePermission := range rolePermissionList {
		permissionIDList[i] = rolePermission.PermissionID
	}
	permissionList, _, err := p.svc.Dao.PermissionDao(ctx).WhereIDIn(permissionIDList).Find()

	permissionPaths := make([]string, len(permissionList))
	for i, permission := range permissionList {
		permissionPaths[i] = permission.Path
	}

	return &api.GetUserPermissionPathResponse{
		Data: &api.GetUserPermissionPathResponseData{
			Permissions: permissionPaths,
		},
	}, nil
}
