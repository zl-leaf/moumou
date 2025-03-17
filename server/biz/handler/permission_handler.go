package handler

import (
	"context"

	"github.com/moumou/server/biz/model"

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

func (p PermissionHandler) DeletePermission(ctx context.Context, request *api.DeletePermissionRequest) (*api.DeletePermissionResponse, error) {
	//TODO implement me
	panic("implement me")
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
	permissionList, _, err := p.svc.Dao.PermissionDao(ctx).OrderBySort().Find()
	if err != nil {
		return nil, err
	}
	treeNodeList := p.converter.ConvertPermissionTreeNodeListToVO(permissionList)
	// 创建mapByID
	treeNodeMapByID := make(map[int64]*api.PermissionTreeNode, len(treeNodeList))
	for _, treeNode := range treeNodeList {
		treeNodeMapByID[treeNode.Id] = treeNode
	}
	root := make([]*api.PermissionTreeNode, 0, len(treeNodeList)/3)
	for _, treeNode := range treeNodeList {
		if treeNode.Pid == 0 {
			// 根节点
			root = append(root, treeNode)
			continue
		}

		pid := treeNode.Pid
		treeNodeMapByID[pid].Children = append(treeNodeMapByID[pid].Children, treeNode)
	}
	return &api.GetPermissionTreeResponse{
		Data: &api.GetPermissionTreeResponseData{
			List: root,
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
	permissionList, _, err := p.svc.Dao.PermissionDao(ctx).WhereIdIn(permissionIDList).Find()

	permissionPaths := make([]string, len(permissionList))
	for i, permission := range permissionList {
		permissionPaths[i] = permission.Code
	}

	return &api.GetUserPermissionPathResponse{
		Data: &api.GetUserPermissionPathResponseData{
			Permissions: permissionPaths,
		},
	}, nil
}
