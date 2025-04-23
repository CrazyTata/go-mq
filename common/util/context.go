package util

import (
	"context"

	"go-mq/common/xerr"
)

// GetUserIdFromContext 从上下文中提取 user_id
func GetUserIdFromContext(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return "", xerr.NewErrCode(xerr.LoginMiss)
	}
	return userId, nil
}

// contextKey 定义上下文键类型
type contextKey string

const (
	// UserIDKey 用于存储 user_id 的上下文键
	UserIDKey contextKey = "user_id"

	// UsernameKey 用于存储 username 的上下文键
	UsernameKey contextKey = "username"
)

// GetUsernameFromContext 从上下文中提取 username
func GetUsernameFromContext(ctx context.Context) (string, error) {
	username, ok := ctx.Value(UsernameKey).(string)
	if !ok {
		return "", xerr.NewErrCode(xerr.LoginMiss)
	}
	return username, nil
}

// GetAppIdFromContext 从上下文中提取 app_id
func GetAppIdFromContext(ctx context.Context) (string, error) {
	return "", nil
}
