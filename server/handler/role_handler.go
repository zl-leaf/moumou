package handler

import (
	"context"
	"github.com/moumou/server/biz/service"
	api "github.com/moumou/server/proto_gen"
)

type RoleHandler struct {
	svc *service.Service
}

func NewRoleHandler(svc *service.Service) api.RoleHandlerHTTPServer {
	return &RoleHandler{}
}

func (r RoleHandler) CreateRole(ctx context.Context, request *api.CreateRoleRequest) (*api.CreateRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) DeleteRole(ctx context.Context, request *api.DeleteRoleRequest) (*api.DeleteRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) UpdateRole(ctx context.Context, request *api.UpdateRoleRequest) (*api.UpdateRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RoleHandler) GetRoleList(ctx context.Context, request *api.GetRoleListRequest) (*api.GetRoleListResponse, error) {
	//TODO implement me
	panic("implement me")
}
