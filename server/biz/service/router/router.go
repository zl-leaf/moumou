package router

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/router/internal"
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

func (svc *Service) GetRouterInfo(ctx context.Context, id int) (*model.Router, error) {
	return svc.manageSvc.GetByID(ctx, id)
}
