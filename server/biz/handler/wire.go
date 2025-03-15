package handler

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/conv/factory"
)

var ProviderSet = wire.NewSet(
	NewUserHandler,
	NewRoleHandler,
	NewSecurityHandler,
	NewPermissionHandler,

	factory.NewConverter,
)
