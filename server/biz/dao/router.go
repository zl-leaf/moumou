package dao

import (
	"context"
	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
)

type routerDao struct {
	*gorm.DB
}

func newRouterDao(db *gorm.DB) *routerDao {
	return &routerDao{db}
}

func (d *routerDao) WithContext(ctx context.Context) *routerDao {
	d.DB = d.DB.WithContext(ctx)
	return d
}

func (d *routerDao) OrderBySort() *routerDao {
	d.DB = d.DB.Order("sort asc")
	return d
}

func (d *routerDao) GetByID(id int64) (*model.Router, error) {
	var record *model.Router

	var query = d.First(&record, id)
	if err := query.Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (d *routerDao) Find() ([]*model.Router, int64, error) {
	var list []*model.Router
	var total int64
	var query = d.Model(model.Router{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *routerDao) Create(record *model.Router) error {
	return d.DB.Create(record).Error
}

func (d *routerDao) Save(record *model.Router) error {
	return d.DB.Save(record).Error
}

func (d *routerDao) Delete(conds ...interface{}) error {
	return d.DB.Delete(&model.Router{}, conds...).Error
}
