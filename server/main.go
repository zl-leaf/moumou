package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/service"
	"github.com/moumou/server/biz/service/user/param"
	"github.com/moumou/server/handler"
	"github.com/moumou/server/handler/mw"
	"github.com/moumou/server/pkgs/config"
	api "github.com/moumou/server/proto_gen"
	"go.opentelemetry.io/otel/sdk/trace"
	"os"
)

type serverConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func NewHTTPServer(logger log.Logger, cnf serverConfig, securityCnf param.SecurityConfig) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(mw.CorsFilter),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(tracing.WithTracerProvider(trace.NewTracerProvider())),
			mw.HttpTrace(),
			selector.Server(jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(securityCnf.JWTKey), nil
			})).Match(func(ctx context.Context, operation string) bool {
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
	if cnf.Host != "" || cnf.Port != "" {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%s", cnf.Host, cnf.Port)))
	}
	srv := http.NewServer(opts...)

	var svc, err = service.InitService()
	if err != nil {
		panic(err)
	}
	routerHandler := handler.NewRouterHandler(svc)
	userHandler := handler.NewUserHandler(svc)
	securityHandler := handler.NewSecurityHandler(svc)
	api.RegisterRouterHandlerHTTPServer(srv, routerHandler)
	api.RegisterUserHandlerHTTPServer(srv, userHandler)
	api.RegisterSecurityHandlerHTTPServer(srv, securityHandler)

	return srv
}

func main() {
	var err error

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	err = config.Init("config")
	if err != nil {
		panic(err)
	}

	var serverCnf serverConfig
	var securityCnf param.SecurityConfig
	err = config.GetConfigs([]string{"server", securityCnf.GetConfigName()}, []any{&serverCnf, &securityCnf})
	if err != nil {
		panic(err)
	}

	app := newApp(logger, NewHTTPServer(logger, serverCnf, securityCnf))
	if err := app.Run(); err != nil {
		panic(err)
	}
}
