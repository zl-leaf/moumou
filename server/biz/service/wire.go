//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/dao"
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/router"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/pkgs/database"
)

var serviceSet = wire.NewSet(
	NewService,
	user.NewUserService,
	router.NewRouterService,
	page.NewPageService,

	dao.NewDao,

	database.NewMysqlGorm,
)

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
