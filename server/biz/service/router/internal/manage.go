package internal

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/router/param"
	"gorm.io/gorm"
)

type ManageService struct {
	db *gorm.DB
}

func NewManageService(db *gorm.DB) *ManageService {
	return &ManageService{db: db}
}

func (svc *ManageService) List(ctx context.Context) ([]*model.Router, error) {
	var routerList []*model.Router

	ret := svc.db.Order("sort ASC").Find(&routerList)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return routerList, nil
}

func (svc *ManageService) GetByID(ctx context.Context, id int64) (*model.Router, error) {
	var routerInfo *model.Router

	var query = svc.db.First(&routerInfo, id)
	if err := query.Error; err != nil {
		return nil, err
	}
	return routerInfo, nil
}

func (svc *ManageService) Create(ctx context.Context, data *param.RouterFormData) (int64, error) {
	router := data.Router
	result := svc.db.Create(&router)
	if result.Error != nil {
		return 0, result.Error
	}
	return int64(router.ID), nil
}

func (svc *ManageService) Update(ctx context.Context, data *param.RouterFormData) error {
	router := data.Router
	result := svc.db.Save(&router)
	return result.Error
}

func (svc *ManageService) Delete(ctx context.Context, ids []int64) error {
	result := svc.db.Delete(&model.Router{}, ids)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
