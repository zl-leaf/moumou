package page

import (
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/page/internal"
	"gorm.io/gorm"
)

type Service struct {
	routerService   *internal.RouterService
	pageService     *internal.PageService
	pageDataService *internal.PageDataService
}

func NewPageService(db *gorm.DB) *Service {
	return &Service{
		routerService:   internal.NewRouterService(db),
		pageService:     internal.NewPageService(),
		pageDataService: internal.NewPageDataService(),
	}
}

func (svc *Service) GetRouterList() ([]*model.Router, error) {
	return svc.routerService.GetRouterList()
}

func (svc *Service) GetPage(pageID int64) (*model.Page, error) {
	return svc.pageService.GetPage(pageID)
}

func (svc *Service) GetDataList(pageID int64, currentPage, pageSize int) ([]map[string]any, int64, error) {
	return svc.pageDataService.GetList(pageID, currentPage, pageSize)
}
