package service

import (
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/user"
)

type Service struct {
	UserService *user.Service
	PageService *page.Service
}

func NewService(
	userService *user.Service,
	pageService *page.Service,
) *Service {
	return &Service{
		UserService: userService,
		PageService: pageService,
	}
}
