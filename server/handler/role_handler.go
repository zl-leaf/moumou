package handler

import (
	"context"

	"github.com/moumou/server/biz/model"

	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type RoleHandler struct {
	svc *service.Service
}

func NewRoleHandler(svc *service.Service) api.RoleHandlerHTTPServer {
	return &RoleHandler{svc: svc}
}

func (r RoleHandler) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.CreateRoleResponse, error) {
	role := ConvCreateRequestData2Role(request.Role)
	err := r.svc.Dao.RoleDao.WithContext(ctx).Create(role)
	if err != nil {
		return nil, err
	}
	return &api.CreateRoleResponse{
		Data: &api.CreateRoleResponseData{
			Id: role.ID,
		},
	}, nil
}

func (r RoleHandler) DeleteRole(ctx context.Context, request *api.DeleteRoleRequest) (*api.DeleteRoleResponse, error) {
	err := r.svc.Dao.RoleDao.WithContext(ctx).Delete(request.Ids)
	if err != nil {
		return nil, err
	}
	return &api.DeleteRoleResponse{}, nil
}

func (r RoleHandler) UpdateRole(ctx context.Context, request *api.UpdateRoleRequest) (*api.UpdateRoleResponse, error) {
	role, err := r.svc.Dao.RoleDao.WithContext(ctx).GetByID(request.Role.GetId())
	if err != nil {
		return nil, err
	}
	role.Name = request.Role.Name
	err = r.svc.Dao.RoleDao.WithContext(ctx).Save(role)
	if err != nil {
		return nil, err
	}
	return &api.UpdateRoleResponse{}, nil
}

func (r RoleHandler) GetRoleList(ctx context.Context, request *api.GetRoleListRequest) (*api.GetRoleListResponse, error) {
	roleList, total, err := r.svc.Dao.RoleDao.WithContext(ctx).Page(int(request.CurrentPage), int(request.PageSize)).Find()
	if err != nil {
		return nil, err
	}

	return &api.GetRoleListResponse{
		Data: &api.GetRoleListResponseData{
			List:  ConvRoleList2VO(roleList),
			Total: total,
		},
	}, nil
}

func (r RoleHandler) GetRoleInfo(ctx context.Context, request *api.GetRoleInfoRequest) (*api.GetRoleInfoResponse, error) {
	role, err := r.svc.Dao.RoleDao.WithContext(ctx).GetByID(request.GetId())
	if err != nil {
		return nil, err
	}
	var permissionList []*model.Permission
	if request.GetField() != nil && request.GetField().GetPermission() {
		rolePermissionList, _, err := r.svc.Dao.RolePermissionDao.WithContext(ctx).WhereRoleIDEq(request.GetId()).Find()
		if err != nil {
			return nil, err
		}

		permissionIDs := make([]int64, len(rolePermissionList))
		for i, rolePermission := range rolePermissionList {
			permissionIDs[i] = rolePermission.PermissionID
		}
		if len(permissionIDs) > 0 {
			permissionList, _, err = r.svc.Dao.PermissionDao.WithContext(ctx).WhereIDIn(permissionIDs).Find()
		}
	}

	return &api.GetRoleInfoResponse{
		Data: ConvRole2VO(role, permissionList),
	}, nil
}

func (r RoleHandler) UpdateRolePermission(ctx context.Context, request *api.UpdateRolePermissionRequest) (*api.UpdateRolePermissionResponse, error) {
	err := r.svc.RoleService.UpdateRolePermission(ctx, request.GetId(), request.GetPermissionIds())
	if err != nil {
		return nil, err
	}
	return &api.UpdateRolePermissionResponse{}, nil
}
