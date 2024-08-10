package service

import (
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/router"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
)

type Service struct {
	SysUserService   *user.Service
	SysRouterService *router.Service
	RoleService      *role.Service
	PageService      *page.Service
	Dao              *dao.Dao
}

func NewService(
	sysUserService *user.Service,
	sysRouterService *router.Service,
	roleService *role.Service,
	pageService *page.Service,
	db *dao.Dao,
) *Service {
	return &Service{
		SysUserService:   sysUserService,
		SysRouterService: sysRouterService,
		RoleService:      roleService,
		PageService:      pageService,
		Dao:              db,
	}
}
