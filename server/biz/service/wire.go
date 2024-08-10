//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/service/page"
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/router"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/database"
)

var serviceSet = wire.NewSet(
	NewService,
	user.NewUserService,
	router.NewRouterService,
	role.NewService,
	page.NewPageService,

	dao.NewDao,

	database.NewMysqlGorm,
)

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
