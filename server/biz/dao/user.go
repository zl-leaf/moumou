package dao

import (
	"context"
	"github.com/moumou/server/biz/model"
	"gorm.io/gorm"
)

type userDao struct {
	*gorm.DB
}

func newUserDao(db *gorm.DB) *userDao {
	return &userDao{db}
}

func (d *userDao) WithContext(ctx context.Context) *userDao {
	d.DB = d.DB.WithContext(ctx)
	return d
}

func (d *userDao) WhereUsernameEq(username string) *userDao {
	d.DB = d.DB.Where("username = ?", username)
	return d
}

func (d *userDao) Limit(limit int) *userDao {
	d.DB = d.DB.Limit(limit)
	return d
}

func (d *userDao) Offset(offset int) *userDao {
	d.DB = d.DB.Offset(offset)
	return d
}

func (d *userDao) GetByID(id int64) (*model.User, error) {
	return d.First(id)
}

func (d *userDao) Find() ([]*model.User, int64, error) {
	var list []*model.User
	var total int64
	var query = d.Model(model.User{})

	if err := query.Find(&list).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (d *userDao) First(conds ...interface{}) (*model.User, error) {
	var record = &model.User{}
	result := d.DB.First(record, conds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return record, nil
}
