package internal

import (
	"context"
	"github.com/moumou/server/biz/model"
	"github.com/moumou/server/biz/service/sys_user/param"
	"gorm.io/gorm"
)

type ManageService struct {
	db *gorm.DB
}

func NewManageService(db *gorm.DB) *ManageService {
	return &ManageService{db: db}
}

// List 系统用户列表
func (s *ManageService) List(ctx context.Context, filter *param.SysUserFilter, currentPage, pageSize int) ([]*model.SysUser, int64, error) {
	var userList []*model.SysUser
	var total int64
	var query = s.db.Model(model.SysUser{})

	if err := query.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&userList).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return userList, total, nil
}

// GetByID 通过id查询
func (s *ManageService) GetByID(ctx context.Context, id int) (*model.SysUser, error) {
	var userInfo *model.SysUser
	var query = s.db.First(&userInfo, id)
	if err := query.Error; err != nil {
		return nil, err
	}
	return userInfo, nil
}
