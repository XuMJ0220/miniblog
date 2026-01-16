package contextx

import "context"

// 自定义用于上下文的键.
type (
	// requestIDKey 定义请求 ID 的上下文键.
	requestIDKey struct{}
	// usernameKey 定义用户名的上下文键.
	userNameKey struct{}
)

// WithRequestID 将请求 ID 存放到上下文中.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID 从上下文中提取请求 ID.
func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}

func Username(ctx context.Context) string{
	username,_:=ctx.Value(userNameKey{}).(string)
	return username
}
