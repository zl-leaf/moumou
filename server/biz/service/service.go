package service

import (
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	UserService *user.Service
	RoleService *role.Service
	Dao         *dao.Dao
}

func NewService(
	userService *user.Service,
	roleService *role.Service,
	db *dao.Dao,
) *Service {
	return &Service{
		UserService: userService,
		RoleService: roleService,
		Dao:         db,
	}
}
