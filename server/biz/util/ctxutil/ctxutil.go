package ctxutil

import "context"

const ctxKeyUserid = "ctx_user_id"

func SetCtxUserID(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, ctxKeyUserid, userId)
}

func GetCtxUserID(ctx context.Context) int64 {
	value := ctx.Value(ctxKeyUserid)
	if intValue, ok := value.(int64); ok {
		return intValue
	}
	return 0
}
