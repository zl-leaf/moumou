package internal

import (
	"context"

	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/gen/dao"
)

type BindUserService struct {
	db *dao.Dao
}

func NewBindUserService(db *dao.Dao) *BindUserService {
	return &BindUserService{db: db}
}

func (s *BindUserService) Execute(ctx context.Context, roleID int64, bindUserIDList []int64) error {
	return s.db.Transaction(func(tx *dao.Dao) error {
		existsBindUserList, _, err := tx.UserRelRoleDao.WithContext(ctx).WhereRoleIDEq(roleID).Find()
		if err != nil {
			return err
		}
		var (
			needBindUserIDs = make([]int64, 0, len(bindUserIDList))
			unBinduserIDs   = make([]int64, 0, len(existsBindUserList))
		)

		// 确定需要添加的绑定的userID
		existsBindUserIDMap := make(map[int64]bool, len(existsBindUserList))
		for _, bindUser := range existsBindUserList {
			existsBindUserIDMap[bindUser.UserID] = true
		}
		for _, bindUserID := range bindUserIDList {
			if existsBindUserIDMap[bindUserID] {
				continue
			}
			needBindUserIDs = append(needBindUserIDs, bindUserID)
		}

		// 确定需要接触绑定的userID
		bindUserIDMap := make(map[int64]bool, len(bindUserIDList))
		for _, bindUserID := range bindUserIDList {
			bindUserIDMap[bindUserID] = true
		}
		for _, bindUser := range existsBindUserList {
			if !bindUserIDMap[bindUser.UserID] {
				unBinduserIDs = append(unBinduserIDs, bindUser.UserID)
			}
		}

		if len(unBinduserIDs) > 0 {
			if err = tx.UserRelRoleDao.WithContext(ctx).Delete(unBinduserIDs); err != nil {
				return err
			}
		}

		if len(needBindUserIDs) > 0 {
			for _, needBindUserID := range needBindUserIDs {
				err = tx.UserRelRoleDao.WithContext(ctx).Create(&model.UserRelRole{
					RoleID: roleID,
					UserID: needBindUserID,
				})
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}
