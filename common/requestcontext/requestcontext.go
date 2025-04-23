package requestcontext

import "context"

// contextKey 定义上下文键类型
type contextKey string

const (
	// AppSecretKey 用于存储 appSecret 的上下文键
	AppSecretKey contextKey = "app_secret"
	// MerchantIDKey 用于存储 merchant_id 的上下文键
	MerchantIDKey contextKey = "merchant_id"
)

// GetAppSecretFromContext 安全地从上下文中获取 appSecret
func GetAppSecretFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	value := ctx.Value(AppSecretKey)
	if value == nil {
		return "", false
	}
	secret, ok := value.(string)
	return secret, ok
}

// GetMerchantIDFromContext 安全地从上下文中获取 merchant_id
func GetMerchantIDFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	value := ctx.Value(MerchantIDKey)
	if value == nil {
		return "", false
	}
	id, ok := value.(string)
	return id, ok
}
