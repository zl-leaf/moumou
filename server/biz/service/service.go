package service

import (
	"github.com/moumou/server/biz/service/permission"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	UserService       *user.Service
	PermissionService *permission.Service
	Dao               *dao.Dao
}

func NewService(
	userService *user.Service,
	permissionService *permission.Service,
	db *dao.Dao,
) *Service {
	return &Service{
		UserService:       userService,
		PermissionService: permissionService,
		Dao:               db,
	}
}
