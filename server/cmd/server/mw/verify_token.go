package mw

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/moumou/server/biz/conf"
	"github.com/moumou/server/biz/service/user"
	userdata "github.com/moumou/server/biz/service/user/data"
	"github.com/moumou/server/biz/util/ctxutil"
)

// JWTServer JWT中间件
func JWTServer(confData *conf.Data) middleware.Middleware {
	return jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
		return []byte(confData.SecurityConfig.JWTKey), nil
	}, jwt.WithClaims(func() jwtv5.Claims {
		return &userdata.CustomClaims{}
	}))
}

// VerifyUser 验证用户
func VerifyUser(userService *user.Service) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			userId, err := userService.VerifyToken(ctx)
			if err != nil {
				return nil, errors.Unauthorized("Unauthorized", "非法用户")
			}
			ctx = ctxutil.SetCtxUserID(ctx, userId)
			return handler(ctx, req)
		}
	}
}
