package router

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/router/internal"
	"github.com/moumou/server/biz/service/router/param"
	"gorm.io/gorm"
)

type Service struct {
	manageSvc *internal.ManageService
}

func NewRouterService(db *gorm.DB) *Service {
	return &Service{
		manageSvc: internal.NewManageService(db),
	}
}

func (svc *Service) GetRouterList(ctx context.Context) ([]*model.Router, error) {
	return svc.manageSvc.List(ctx)
}

func (svc *Service) GetRouterInfo(ctx context.Context, id int64) (*model.Router, error) {
	return svc.manageSvc.GetByID(ctx, id)
}

func (svc *Service) CreateRouter(ctx context.Context, data *param.RouterFormData) (int64, error) {
	return svc.manageSvc.Create(ctx, data)
}

func (svc *Service) UpdateRouter(ctx context.Context, data *param.RouterFormData) error {
	return svc.manageSvc.Update(ctx, data)
}

func (svc *Service) DeleteRouter(ctx context.Context, ids []int64) error {
	// 删除节点与其下子节点
	allRouter, err := svc.GetRouterList(ctx)
	if err != nil {
		return err
	}

	// 删除其下所有子节点
	var parentIdMap = make(map[int64]bool, 2*len(ids))
	for _, id := range ids {
		parentIdMap[id] = true
	}
	for len(parentIdMap) != 0 {
		needDeleteIds := make([]int64, 0, len(allRouter))
		newParentIdMap := make(map[int64]bool, len(parentIdMap))
		for _, router := range allRouter {
			if parentIdMap[int64(router.Pid)] {
				needDeleteIds = append(needDeleteIds, router.ID)
				newParentIdMap[router.ID] = true
			}
		}
		parentIdMap = newParentIdMap

		if len(needDeleteIds) > 0 {
			err := svc.manageSvc.Delete(ctx, needDeleteIds)
			if err != nil {
				return err
			}
		}
	}
	return svc.manageSvc.Delete(ctx, ids)
}
