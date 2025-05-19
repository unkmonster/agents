package biz

import "context"

type GatewayRepo interface {
	// CreateUserConsumer  创建消费者并返回消费者标识 (id or username)
	CreateUserConsumer(ctx context.Context, user *User) (string, error)
	// EnableJwtPluginForConsumer 启用 JWT 插件, 返回 jwt key id
	EnableJwtPluginForConsumer(ctx context.Context, consumer string, pubkeyPem string) (string, error)
}
