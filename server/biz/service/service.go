package service

import (
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	SysUserService *user.Service
	RoleService    *role.Service
	Dao            *dao.Dao
}

func NewService(
	sysUserService *user.Service,
	roleService *role.Service,
	db *dao.Dao,
) *Service {
	return &Service{
		SysUserService: sysUserService,
		RoleService:    roleService,
		Dao:            db,
	}
}
