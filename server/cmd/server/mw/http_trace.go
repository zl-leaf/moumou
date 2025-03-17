package mw

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
)

func HttpTrace() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				defer func() {
					tr.ReplyHeader().Set("Span-Id", tracing.SpanID()(ctx).(string))
					tr.ReplyHeader().Set("Trace-Id", tracing.TraceID()(ctx).(string))
				}()
			}
			return handler(ctx, req)
		}
	}
}
