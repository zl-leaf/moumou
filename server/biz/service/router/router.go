package router

import (
	"context"
	"github.com/moumou/server/biz/service/router/internal"
	"github.com/moumou/server/biz/service/router/param"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	manageSvc *internal.ManageService
	db        *dao.Dao
}

func NewRouterService(db *dao.Dao) *Service {
	return &Service{
		manageSvc: internal.NewManageService(db),
		db:        db,
	}
}

func (svc *Service) CreateRouter(ctx context.Context, data *param.RouterFormData) (int64, error) {
	router := data.Router
	err := svc.db.RouterDao.WithContext(ctx).Create(&router)
	return router.ID, err
}

func (svc *Service) UpdateRouter(ctx context.Context, data *param.RouterFormData) error {
	return svc.db.RouterDao.WithContext(ctx).Save(&data.Router)
}

func (svc *Service) DeleteRouter(ctx context.Context, ids []int64) error {
	// 删除节点与其下子节点
	return svc.manageSvc.DeleteWithChildren(ctx, ids)
}
