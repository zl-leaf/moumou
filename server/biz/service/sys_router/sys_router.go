package sys_router

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/sys_router/internal"
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

func (svc *Service) GetSysRouterList(ctx context.Context) ([]*model.SysRouter, error) {
	return svc.manageSvc.List(ctx)
}

func (svc *Service) GetSysRouterInfo(ctx context.Context, id int) (*model.SysRouter, error) {
	return svc.manageSvc.GetByID(ctx, id)
}
