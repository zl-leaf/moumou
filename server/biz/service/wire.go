package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/service/permission"
	"github.com/moumou/server/biz/service/system"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/gen/dao"
	"github.com/moumou/server/pkgs/database"
)

var ProviderSet = wire.NewSet(
	NewService,
	user.NewUserService,
	permission.NewService,
	system.NewService,

	dao.NewDao,

	database.NewMemoryGorm,
)
