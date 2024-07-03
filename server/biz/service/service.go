package service

import (
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/sys_router"
	"github.com/moumou/server/biz/service/sys_user"
)

type Service struct {
	SysUserService   *sys_user.Service
	SysRouterService *sys_router.Service
	PageService      *page.Service
}

func NewService(
	sysUserService *sys_user.Service,
	sysRouterService *sys_router.Service,
	pageService *page.Service,
) *Service {
	return &Service{
		SysUserService:   sysUserService,
		SysRouterService: sysRouterService,
		PageService:      pageService,
	}
}
