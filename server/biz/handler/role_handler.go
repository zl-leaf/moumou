package handler

import (
	"context"

	"github.com/moumou/server/biz/conv"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/gen/proto"
)

type RoleHandler struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewRoleHandler(svc *service.Service, converter conv.IConverter) api.RoleHandlerHTTPServer {
	return &RoleHandler{svc: svc, converter: converter}
}

func (r RoleHandler) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.CreateRoleResponse, error) {
	role := new(model.Role)
	r.converter.ConvertCreateRoleRequestDataToBO(request.GetRole(), role)
	err := r.svc.Dao.RoleDao(ctx).Create(role)
	if err != nil {
		return nil, err
	}
	return &api.CreateRoleResponse{
		Data: &api.CreateRoleResponseData{
			Id: role.Id,
		},
	}, nil
}

func (r RoleHandler) DeleteRole(ctx context.Context, request *api.DeleteRoleRequest) (*api.DeleteRoleResponse, error) {
	err := r.svc.Dao.RoleDao(ctx).Delete(request.Ids)
	if err != nil {
		return nil, err
	}
	return &api.DeleteRoleResponse{}, nil
}

func (r RoleHandler) UpdateRole(ctx context.Context, request *api.UpdateRoleRequest) (*api.UpdateRoleResponse, error) {
	role, err := r.svc.Dao.RoleDao(ctx).GetByID(request.Role.GetId())
	if err != nil {
		return nil, err
	}
	r.converter.ConvertUpdateRoleRequestDataToBO(request.GetRole(), role)
	err = r.svc.Dao.RoleDao(ctx).Save(role)
	if err != nil {
		return nil, err
	}
	return &api.UpdateRoleResponse{}, nil
}

func (r RoleHandler) GetRoleList(ctx context.Context, request *api.GetRoleListRequest) (*api.GetRoleListResponse, error) {
	roleList, total, err := r.svc.Dao.RoleDao(ctx).Page(int(request.CurrentPage), int(request.PageSize)).Find()
	if err != nil {
		return nil, err
	}

	return &api.GetRoleListResponse{
		Data: &api.GetRoleListResponseData{
			List:  r.converter.ConvertRoleListToVO(roleList),
			Total: total,
		},
	}, nil
}

func (r RoleHandler) GetRoleInfo(ctx context.Context, request *api.GetRoleInfoRequest) (*api.GetRoleInfoResponse, error) {
	role, err := r.svc.Dao.RoleDao(ctx).GetByID(request.GetId())
	if err != nil {
		return nil, err
	}
	var permissionList []*model.Permission
	if request.GetField() != nil && request.GetField().GetPermission() {
		permissionList, err = r.svc.RoleService.GetPermissionsByRoleId(ctx, request.GetId())
		if err != nil {
			return nil, err
		}
	}

	var userList []*model.User
	if request.GetField() != nil && request.GetField().GetBindUser() {
		userList, err = r.svc.RoleService.GetBindUserByRoleId(ctx, request.GetId())
		if err != nil {
			return nil, err
		}
	}

	roleVO := r.converter.ConvertRoleToVO(role)
	roleVO.Permissions = r.converter.ConvertPermissionListToVO(permissionList)
	roleVO.Users = r.converter.ConvertUserListToVO(userList)

	return &api.GetRoleInfoResponse{
		Data: roleVO,
	}, nil
}

func (r RoleHandler) UpdateRolePermission(ctx context.Context, request *api.UpdateRolePermissionRequest) (*api.UpdateRolePermissionResponse, error) {
	err := r.svc.RoleService.UpdateRolePermission(ctx, request.GetId(), request.GetPermissionIds())
	if err != nil {
		return nil, err
	}
	return &api.UpdateRolePermissionResponse{}, nil
}

func (r RoleHandler) UpdateBindUser(ctx context.Context, request *api.UpdateBindUserRequest) (*api.UpdateBindUserResponse, error) {
	err := r.svc.RoleService.BindUsers(ctx, request.GetRoleId(), request.GetUserIds())
	if err != nil {
		return nil, err
	}
	return &api.UpdateBindUserResponse{}, nil
}
