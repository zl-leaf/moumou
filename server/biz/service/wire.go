//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/sys_router"
	"github.com/moumou/server/biz/service/sys_user"
	"github.com/moumou/server/framework/database"
)

var serviceSet = wire.NewSet(
	NewService,
	sys_user.NewSysUserService,
	sys_router.NewRouterService,
	page.NewPageService,
	database.NewMysqlGorm,
)

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
