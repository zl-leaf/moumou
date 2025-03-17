package role

import (
	"context"

	"github.com/moumou/server/biz/service/role/internal"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	db *dao.Dao
}

func NewService(db *dao.Dao) *Service {
	return &Service{db: db}
}

func (s *Service) UpdateRolePermission(ctx context.Context, roleID int64, permissionIDs []int64) error {
	return internal.NewUpdateRolePermissionService(s.db).Execute(ctx, roleID, permissionIDs)
}

func (s *Service) BindUsers(ctx context.Context, roleID int64, userIDs []int64) error {
	return internal.NewBindUserService(s.db).Execute(ctx, roleID, userIDs)
}

func (s *Service) GetPermissionCodesByUid(ctx context.Context, userId int64) ([]string, error) {
	userRoleList, _, err := s.db.UserRelRoleDao(ctx).WhereUserIDEq(userId).Find()
	if err != nil {
		return nil, err
	}
	roleIDs := make([]int64, len(userRoleList))
	for i, userRole := range userRoleList {
		roleIDs[i] = userRole.RoleID
	}
	rolePermissionList, _, err := s.db.RolePermissionDao(ctx).WhereRoleIDIn(roleIDs).Find()
	if err != nil {
		return nil, err
	}
	permissionIDList := make([]int64, len(rolePermissionList))
	for i, rolePermission := range rolePermissionList {
		permissionIDList[i] = rolePermission.PermissionID
	}
	permissionList, _, err := s.db.PermissionDao(ctx).WhereIdIn(permissionIDList).Find()
	if err != nil {
		return nil, err
	}

	permissionCodes := make([]string, len(permissionList))
	for i, permission := range permissionList {
		permissionCodes[i] = permission.Code
	}
	return permissionCodes, nil
}
