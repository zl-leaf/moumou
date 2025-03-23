package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/cmd/server/mw"
	api "github.com/moumou/server/gen/proto"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewHTTPServer(
	logger log.Logger,
	confData *conf.Data,
	userHandler api.UserHandlerHTTPServer,
	roleHandler api.RoleHandlerHTTPServer,
	securityHandler api.SecurityHandlerHTTPServer,
	permissionHandler api.PermissionHandlerHTTPServer,
	svc *service.Service,
) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(mw.CorsFilter),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(trace.NewTracerProvider())),
			mw.HttpTrace(),
			selector.Server(
				mw.JWTServer(confData),
				mw.VerifyUser(svc.UserService)).Match(func(ctx context.Context, operation string) bool {
				whiteList := []string{api.OperationSecurityHandlerLogin}
				for _, white := range whiteList {
					if operation == white {
						return false
					}
				}
				return true
			}).Build(),
			logging.Server(logger),
			mw.APIPermission(svc),
		),
	}
	if confData.ServerConfig.Host != "" || confData.ServerConfig.Port != "" {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%s", confData.ServerConfig.Host, confData.ServerConfig.Port)))
	}
	if confData.ServerConfig.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Second*time.Duration(confData.ServerConfig.Timeout)))
	}
	srv := http.NewServer(opts...)

	api.RegisterUserHandlerHTTPServer(srv, userHandler)
	api.RegisterRoleHandlerHTTPServer(srv, roleHandler)
	api.RegisterSecurityHandlerHTTPServer(srv, securityHandler)
	api.RegisterPermissionHandlerHTTPServer(srv, permissionHandler)

	return srv
}
