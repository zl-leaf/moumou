package internal

import (
	"context"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/gen/dao"
)

type UpdateRolePermissionService struct {
	db *dao.Dao
}

func NewUpdateRolePermissionService(db *dao.Dao) *UpdateRolePermissionService {
	return &UpdateRolePermissionService{
		db: db,
	}
}

func (s *UpdateRolePermissionService) Execute(ctx context.Context, roleID int64, permissionIDs []int64) error {
	return s.db.Transaction(ctx, func(tx *dao.Dao) error {
		rolePermissionList, _, err := tx.RolePermissionDao(ctx).WhereRoleIDEq(roleID).Find()
		if err != nil {
			return err
		}

		deleteByPermissionIDs := make([]int64, 0, len(rolePermissionList))
		addByPermissionIDs := make([]int64, 0, len(permissionIDs))

		// 整理出需要添加的PermissionID
		existsPermissionIdMap := make(map[int64]bool, len(rolePermissionList))
		for _, rolePermission := range rolePermissionList {
			existsPermissionIdMap[rolePermission.PermissionID] = true
		}
		for _, permissionID := range permissionIDs {
			if !existsPermissionIdMap[permissionID] {
				addByPermissionIDs = append(addByPermissionIDs, permissionID)
			}
		}

		// 整理出需要删除的PermissionID
		newPermissionIdMap := make(map[int64]bool, len(permissionIDs))
		for _, permissionID := range permissionIDs {
			newPermissionIdMap[permissionID] = true
		}
		for _, rolePermission := range rolePermissionList {
			if !newPermissionIdMap[rolePermission.PermissionID] {
				deleteByPermissionIDs = append(deleteByPermissionIDs, rolePermission.PermissionID)
			}
		}

		if len(deleteByPermissionIDs) > 0 {
			err = tx.RolePermissionDao(ctx).
				WhereRoleIDEq(roleID).
				WherePermissionIDIn(deleteByPermissionIDs).
				Delete()
			if err != nil {
				return err
			}
		}
		if len(addByPermissionIDs) > 0 {
			for _, addByPermissionID := range addByPermissionIDs {
				err = tx.RolePermissionDao(ctx).Create(&model.RolePermission{
					RoleID:       roleID,
					PermissionID: addByPermissionID,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
