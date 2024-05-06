package page

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/page/internal"
	"gorm.io/gorm"
)

type Service struct {
	routerService *internal.RouterService
}

func NewPageService(db *gorm.DB) *Service {
	return &Service{
		routerService: internal.NewRouterService(db),
	}
}

func (svc *Service) GetRouterList() ([]*model.Router, error) {
	return svc.routerService.GetRouterList()
}
