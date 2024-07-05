package internal

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/user/param"
	"gorm.io/gorm"
)

type ManageService struct {
	db *gorm.DB
}

func NewManageService(db *gorm.DB) *ManageService {
	return &ManageService{db: db}
}

// List 系统用户列表
func (s *ManageService) List(ctx context.Context, filter *param.ListUserFilter, currentPage, pageSize int) ([]*model.User, int64, error) {
	var userList []*model.User
	var total int64
	var query = s.db.Model(model.User{})

	if err := query.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&userList).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return userList, total, nil
}

// GetByID 通过id查询
func (s *ManageService) GetByID(ctx context.Context, id int) (*model.User, error) {
	var userInfo *model.User
	var query = s.db.First(&userInfo, id)
	if err := query.Error; err != nil {
		return nil, err
	}
	return userInfo, nil
}
