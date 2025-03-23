package role

import (
	"context"

	"github.com/moumou/server/biz/model"

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

func (s *Service) GetTopLevelPermissions(ctx context.Context) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}
	return tree.GetTopLevelPermissions(), nil
}

func (s *Service) GetPermissionCodesByUid(ctx context.Context, userId int64) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}

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
	permissionList, _ := tree.GetPermissionsFullPathByIds(permissionIDList)
	return permissionList, nil
}

func (s *Service) GetPermissionsByRoleId(ctx context.Context, roleId int64) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}

	rolePermissionList, _, err := s.db.RolePermissionDao(ctx).WhereRoleIDEq(roleId).Find()
	if err != nil {
		return nil, err
	}
	permissionIDs := make([]int64, len(rolePermissionList))
	for i, rolePermission := range rolePermissionList {
		permissionIDs[i] = rolePermission.PermissionID
	}

	permissionList, _ := tree.GetPermissionsFullPathByIds(permissionIDs)
	return permissionList, err
}

func (s *Service) GetBindUserByRoleId(ctx context.Context, roleId int64) ([]*model.User, error) {
	userRelRoleList, _, err := s.db.UserRelRoleDao(ctx).WhereRoleIDEq(roleId).Find()
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, len(userRelRoleList))
	for i, userRelRole := range userRelRoleList {
		userIDs[i] = userRelRole.UserID
	}
	var userList []*model.User
	if len(userIDs) > 0 {
		userList, _, err = s.db.UserDao(ctx).WhereIdIn(userIDs).Find()
	}
	return userList, err
}
