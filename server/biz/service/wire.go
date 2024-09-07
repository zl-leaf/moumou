package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/service/role"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/database"
)

var ProviderSet = wire.NewSet(
	NewService,
	user.NewUserService,
	role.NewService,

	dao.NewDao,

	database.NewMysqlGorm,
)
