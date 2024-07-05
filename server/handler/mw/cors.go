package mw

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func CorsMW() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				// TODO 如果是options就返回成功

				defer func() {
					// Do something on exiting
					tr.ReplyHeader().Add("Access-Control-Allow-Origin", "*")
					tr.ReplyHeader().Add("Access-Control-Allow-Method", "*")
					tr.ReplyHeader().Add("Access-Control-Allow-Headers", "*")
					tr.ReplyHeader().Add("Access-Control-Allow-Credentials", "true")
				}()
			}
			return handler(ctx, req)
		}
	}
}
