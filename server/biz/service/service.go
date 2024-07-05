package service

import (
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/router"
	"github.com/moumou/server/biz/service/user"
)

type Service struct {
	SysUserService   *user.Service
	SysRouterService *router.Service
	PageService      *page.Service
}

func NewService(
	sysUserService *user.Service,
	sysRouterService *router.Service,
	pageService *page.Service,
) *Service {
	return &Service{
		SysUserService:   sysUserService,
		SysRouterService: sysRouterService,
		PageService:      pageService,
	}
}
