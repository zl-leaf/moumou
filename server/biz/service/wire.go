//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/moumou/server/biz/service/user"
	"github.com/moumou/server/framework/database"
)

var serviceSet = wire.NewSet(NewService, user.NewUserService, database.NewMysqlGorm)

func InitService() (*Service, error) {
	wire.Build(serviceSet)
	return nil, nil
}
