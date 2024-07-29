package internal

import (
	"context"
	"errors"
	"github.com/moumou/server/biz/dao"
)

type ManageService struct {
	db *dao.Dao
}

func NewManageService(db *dao.Dao) *ManageService {
	return &ManageService{db: db}
}

func (svc *ManageService) DeleteWithChildren(ctx context.Context, ids []int64) error {
	// 删除节点与其下子节点
	allRouter, _, err := svc.db.RouterDao.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return svc.db.Transaction(func(tx *dao.Dao) error {
		var routerDB = tx.RouterDao.WithContext(ctx)
		// 删除其下所有子节点
		var parentIdMap = make(map[int64]bool, 2*len(ids))
		for _, id := range ids {
			parentIdMap[id] = true
		}
		needDeleteIds := make([]int64, 0, len(allRouter))
		needDeleteIds = append(needDeleteIds, ids...)

		for len(parentIdMap) != 0 {
			newParentIdMap := make(map[int64]bool, len(parentIdMap))
			for _, router := range allRouter {
				if parentIdMap[int64(router.Pid)] {
					if router.IsSystem {
						return errors.New("系统路由禁止删除")
					}

					needDeleteIds = append(needDeleteIds, router.ID)
					newParentIdMap[router.ID] = true
				}
			}
			parentIdMap = newParentIdMap

		}
		if len(needDeleteIds) > 0 {
			err := routerDB.Delete(needDeleteIds)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
