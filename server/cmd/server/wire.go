//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/handler"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/pkgs/database"
)

func wireApp(log.Logger, *conf.Data, *database.DbConfig) (*kratos.App, error) {
	panic(wire.Build(service.ProviderSet, handler.ProviderSet, NewHTTPServer, newApp))
}
