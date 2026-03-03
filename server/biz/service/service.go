package service

import (
	"github.com/moumou/server/biz/service/permission"
	"github.com/moumou/server/biz/service/system"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	UserService       *user.Service
	PermissionService *permission.Service
	SystemService     *system.Service
	Dao               *dao.Dao
}

func NewService(
	userService *user.Service,
	permissionService *permission.Service,
	systemService *system.Service,
	dao *dao.Dao,
) *Service {
	return &Service{
		UserService:       userService,
		PermissionService: permissionService,
		SystemService:     systemService,
		Dao:               dao,
	}
}
