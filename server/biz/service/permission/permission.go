package permission

import (
	"context"

	"github.com/bytedance/gg/gslice"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/permission/internal"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	db *dao.Dao
}

func NewService(db *dao.Dao) *Service {
	return &Service{db: db}
}

// UpdateRolePermission 更新角色与权限绑定关系
func (s *Service) UpdateRolePermission(ctx context.Context, roleID int64, permissionIDs []int64) error {
	rolePermissionList := gslice.Map(permissionIDs, func(permissionID int64) *model.RolePermission {
		return &model.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
		}
	})

	return internal.NewBindRelationService(s.db).BatchUpdateRolePermissionRelation(ctx, rolePermissionList)
}

// BindUsers 角色绑定用户
func (s *Service) BindUsers(ctx context.Context, roleID int64, userIDs []int64) error {
	userRoleList := gslice.Map(userIDs, func(userID int64) *model.UserRelRole {
		return &model.UserRelRole{
			UserID: userID,
			RoleID: roleID,
		}
	})

	return internal.NewBindRelationService(s.db).BatchUpdateUserRoleRelation(ctx, userRoleList)
}

// GetTopLevelPermissions 获取一级权限列表
func (s *Service) GetTopLevelPermissions(ctx context.Context) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}
	return tree.GetTopLevelPermissions(), nil
}

// GetPermissionCodesByUid 获取用户拥有的权限code
func (s *Service) GetPermissionCodesByUid(ctx context.Context, userId int64) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}

	userRoleList, _, err := s.db.UserRelRoleDao(ctx).WhereUserIDEq(userId).Find()
	if err != nil {
		return nil, err
	}
	roleIDs := gslice.Map(userRoleList, func(m *model.UserRelRole) int64 {
		return m.RoleID
	})
	rolePermissionList, _, err := s.db.RolePermissionDao(ctx).WhereRoleIDIn(roleIDs).Find()
	if err != nil {
		return nil, err
	}
	permissionIDList := gslice.Map(rolePermissionList, func(m *model.RolePermission) int64 {
		return m.PermissionID
	})
	permissionList, _ := tree.GetPermissionsFullPathByIds(permissionIDList)
	return permissionList, nil
}

// GetPermissionsByRoleId 获取角色下的权限
func (s *Service) GetPermissionsByRoleId(ctx context.Context, roleId int64, isFullPath bool) ([]*model.Permission, error) {
	tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
	if err != nil {
		return nil, err
	}

	rolePermissionList, _, err := s.db.RolePermissionDao(ctx).WhereRoleIDEq(roleId).Find()
	if err != nil {
		return nil, err
	}
	permissionIDs := gslice.Map(rolePermissionList, func(m *model.RolePermission) int64 {
		return m.PermissionID
	})
	var permissionList []*model.Permission
	if isFullPath {
		permissionList, _ = tree.GetPermissionsFullPathByIds(permissionIDs)
	} else {
		permissionList = make([]*model.Permission, 0, len(permissionIDs))
		for _, permissionID := range permissionIDs {
			if permission := tree.GetPermissionById(permissionID); permission != nil {
				permissionList = append(permissionList, permission)
			}
		}
	}

	return permissionList, err
}

// GetBindUserByRoleId 获取角色绑定的用户
func (s *Service) GetBindUserByRoleId(ctx context.Context, roleId int64) ([]*model.User, error) {
	userRelRoleList, _, err := s.db.UserRelRoleDao(ctx).WhereRoleIDEq(roleId).Find()
	if err != nil {
		return nil, err
	}
	userIDs := gslice.Map(userRelRoleList, func(m *model.UserRelRole) int64 {
		return m.UserID
	})
	var userList []*model.User
	if len(userIDs) > 0 {
		userList, _, err = s.db.UserDao(ctx).WhereIdIn(userIDs).Find()
	}
	return userList, err
}

// DeleteRole 删除角色
func (s *Service) DeleteRole(ctx context.Context, roleIDs []int64) error {
	return s.db.Transaction(ctx, func(tx *dao.Dao) error {
		var err error
		// 解绑user role关系
		err = tx.UserRelRoleDao(ctx).WhereRoleIDIn(roleIDs).Delete()
		if err != nil {
			return err
		}

		// 解绑role permission
		err = tx.RolePermissionDao(ctx).WhereRoleIDIn(roleIDs).Delete()
		if err != nil {
			return err
		}

		// 删除角色
		err = tx.RoleDao(ctx).Delete(roleIDs)
		if err != nil {
			return err
		}

		return nil
	})
}

// DeletePermission 删除权限
func (s *Service) DeletePermission(ctx context.Context, permissionIDs []int64) error {
	return s.db.Transaction(ctx, func(tx *dao.Dao) error {
		var err error
		tree, err := internal.NewPermissionGetterService(s.db).GetPermissionTree(ctx)
		if err != nil {
			return err
		}
		fullChildPermissionList := tree.GetChildPermissionFullPathByIds(permissionIDs)
		if len(fullChildPermissionList) == 0 {
			return nil
		}

		fullChildPermissionIDs := gslice.Map(fullChildPermissionList, func(m *model.Permission) int64 {
			return m.Id
		})

		// 解绑role permission
		err = tx.RolePermissionDao(ctx).WherePermissionIDIn(fullChildPermissionIDs).Delete()
		if err != nil {
			return err
		}

		// 删除权限
		err = tx.PermissionDao(ctx).Delete(fullChildPermissionIDs)
		if err != nil {
			return err
		}

		return nil
	})
}
