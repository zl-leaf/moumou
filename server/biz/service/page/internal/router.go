package internal

import (
	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
)

type RouterService struct {
	db *gorm.DB
}

func NewRouterService(db *gorm.DB) *RouterService {
	return &RouterService{
		db: db,
	}
}

func (svc *RouterService) GetRouterList() ([]*model.SysRouter, error) {
	var routerList []*model.SysRouter
	ret := svc.db.Order("sort ASC").Find(&routerList)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return routerList, nil
}
