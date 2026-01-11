package internal

import (
	"context"
	"fmt"

	"github.com/bytedance/gg/gslice"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/gen/dao"
)

type BindRelationService struct {
	db *dao.Dao
}

func NewBindRelationService(db *dao.Dao) *BindRelationService {
	return &BindRelationService{db: db}
}

// BatchUpdateUserRoleRelation 批量更新用户与角色的绑定关系
func (s *BindRelationService) BatchUpdateUserRoleRelation(ctx context.Context, userRoleList []*model.UserRelRole) error {
	return s.db.Transaction(ctx, func(tx *dao.Dao) error {
		userIDs := gslice.Map(userRoleList, func(m *model.UserRelRole) int64 {
			return m.UserID
		})
		// 获取范围内用户绑定了的角色
		existsBindUserList, _, err := tx.UserRelRoleDao(ctx).WhereUserIDIn(userIDs).Find()
		if err != nil {
			return err
		}

		var (
			needBindUserRefRoleList = make([]*model.UserRelRole, 0, len(userRoleList)) // 需要新增的绑定关系
			unBindIDs               = make([]int64, 0, len(existsBindUserList))        // 需要解绑的关系
		)
		// 获取需要新增的绑定关系
		existsBindUserMap := gslice.ToMap(existsBindUserList, func(m *model.UserRelRole) (string, bool) {
			return fmt.Sprintf("%d_%d", m.UserID, m.RoleID), true
		})
		for _, bindUser := range userRoleList {
			key := fmt.Sprintf("%d_%d", bindUser.UserID, bindUser.RoleID)
			if existsBindUserMap[key] {
				continue
			}
			if bindUser.UserID > 0 && bindUser.RoleID > 0 {
				needBindUserRefRoleList = append(needBindUserRefRoleList, bindUser)
			}
		}

		// 获取需要取消的绑定关系
		userRoleMap := gslice.ToMap(userRoleList, func(m *model.UserRelRole) (string, bool) {
			return fmt.Sprintf("%d_%d", m.UserID, m.RoleID), true
		})
		for _, existsBindUser := range existsBindUserList {
			key := fmt.Sprintf("%d_%d", existsBindUser.UserID, existsBindUser.RoleID)
			if userRoleMap[key] {
				continue
			}
			unBindIDs = append(unBindIDs, existsBindUser.Id)
		}

		if len(unBindIDs) > 0 {
			if err = tx.UserRelRoleDao(ctx).Delete(unBindIDs); err != nil {
				return err
			}
		}

		if len(needBindUserRefRoleList) > 0 {
			for _, needBindUser := range needBindUserRefRoleList {
				err = tx.UserRelRoleDao(ctx).Create(&model.UserRelRole{
					RoleID: needBindUser.RoleID,
					UserID: needBindUser.UserID,
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// BatchUpdateRolePermissionRelation 批量更新角色与权限的关系
func (s *BindRelationService) BatchUpdateRolePermissionRelation(ctx context.Context, rolePermissionList []*model.RolePermission) error {
	return s.db.Transaction(ctx, func(tx *dao.Dao) error {
		roleIDs := gslice.Map(rolePermissionList, func(m *model.RolePermission) int64 {
			return m.RoleID
		})
		existsRolePermissionList, _, err := tx.RolePermissionDao(ctx).WhereRoleIDIn(roleIDs).Find()
		if err != nil {
			return err
		}

		var (
			needBindRolePermissionList = make([]*model.RolePermission, 0, len(rolePermissionList)) // 需要新增的绑定关系
			unBindIDs                  = make([]int64, 0, len(existsRolePermissionList))           // 需要解绑的关系
		)
		// 获取需要新增的绑定关系
		existsBindRolePermissionMap := gslice.ToMap(existsRolePermissionList, func(m *model.RolePermission) (string, bool) {
			return fmt.Sprintf("%d_%d", m.RoleID, m.PermissionID), true
		})
		for _, bindRolePermission := range rolePermissionList {
			key := fmt.Sprintf("%d_%d", bindRolePermission.RoleID, bindRolePermission.PermissionID)
			if existsBindRolePermissionMap[key] {
				continue
			}
			if bindRolePermission.RoleID > 0 && bindRolePermission.PermissionID > 0 {
				needBindRolePermissionList = append(needBindRolePermissionList, bindRolePermission)
			}
		}

		// 获取需要取消的绑定关系
		rolePermissionMap := gslice.ToMap(rolePermissionList, func(m *model.RolePermission) (string, bool) {
			return fmt.Sprintf("%d_%d", m.RoleID, m.PermissionID), true
		})
		for _, existsBindRolePermission := range existsRolePermissionList {
			key := fmt.Sprintf("%d_%d", existsBindRolePermission.RoleID, existsBindRolePermission.PermissionID)
			if rolePermissionMap[key] {
				continue
			}
			unBindIDs = append(unBindIDs, existsBindRolePermission.Id)
		}

		if len(unBindIDs) > 0 {
			if err = tx.RolePermissionDao(ctx).Delete(unBindIDs); err != nil {
				return err
			}
		}

		if len(needBindRolePermissionList) > 0 {
			for _, needBindRolePermission := range needBindRolePermissionList {
				err = tx.RolePermissionDao(ctx).Create(&model.RolePermission{
					RoleID:       needBindRolePermission.RoleID,
					PermissionID: needBindRolePermission.PermissionID,
				})
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
