package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/handler/mw"
	userdata "github.com/moumou/server/biz/service/user/data"
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
) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(mw.CorsFilter),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(trace.NewTracerProvider())),
			mw.HttpTrace(),
			selector.Server(jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(confData.SecurityConfig.JWTKey), nil
			}, jwt.WithClaims(func() jwtv5.Claims {
				return &userdata.CustomClaims{}
			}))).Match(func(ctx context.Context, operation string) bool {
				whiteList := []string{api.OperationSecurityHandlerLogin}
				for _, white := range whiteList {
					if operation == white {
						return false
					}
				}
				return true
			}).Build(),
			logging.Server(logger),
		),
	}
	if confData.ServerConfig.Host != "" || confData.ServerConfig.Port != "" {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%s", confData.ServerConfig.Host, confData.ServerConfig.Port)))
	}
	srv := http.NewServer(opts...)

	//var svc, err = service.InitService(confData, &confData.DBConfig)
	//if err != nil {
	//	panic(err)
	//}
	api.RegisterUserHandlerHTTPServer(srv, userHandler)
	api.RegisterRoleHandlerHTTPServer(srv, roleHandler)
	api.RegisterSecurityHandlerHTTPServer(srv, securityHandler)
	api.RegisterPermissionHandlerHTTPServer(srv, permissionHandler)

	return srv
}
