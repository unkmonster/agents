package biz

import "context"

type CreateConsumerResult struct {
	Consumer string // consumer user_name or id
	Key      string // key for credential
}
type GatewayRepo interface {
	// CreateUserConsumer  创建消费者并返回消费者标识 (id or username)
	CreateUserConsumer(ctx context.Context, user *User) (string, error)
	// EnableJwtPluginForConsumer 启用 JWT 插件, 返回 jwt key id
	CreateConsumerCredential(ctx context.Context, consumer string, pubkeyPem string) (string, error)

	//CreateUserConsumerWithCredential(ctx context.Context, user *User, pubkeyPem string) (res *CreateConsumerResult, err error)
}
