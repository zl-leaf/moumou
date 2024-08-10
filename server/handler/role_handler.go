package handler

import (
	"context"

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
	//TODO implement me
	panic("implement me")
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
	return &api.GetRoleInfoResponse{
		Data: ConvRole2VO(role),
	}, nil
}
