package internal

import (
	"context"
	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
)

type ManageService struct {
	db *gorm.DB
}

func NewManageService(db *gorm.DB) *ManageService {
	return &ManageService{db: db}
}

func (svc *ManageService) List(ctx context.Context) ([]*model.SysRouter, error) {
	var routerList []*model.SysRouter

	ret := svc.db.Order("sort ASC").Find(&routerList)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return routerList, nil
}

func (svc *ManageService) GetByID(ctx context.Context, id int) (*model.SysRouter, error) {
	var routerInfo *model.SysRouter

	var query = svc.db.First(&routerInfo, id)
	if err := query.Error; err != nil {
		return nil, err
	}
	return routerInfo, nil
}
