package middleware

import (
	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

type JwtConfig struct {
	PublicKey string
	Method    string
}

// ServerBasic 返回基础中间件列表
// 1. recovery
// 2. logging
// 3. jwt
// 4. validate
func ServerBasic(logger log.Logger) kratosMiddleware.Middleware {
	return kratosMiddleware.Chain(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		validate.ProtoValidate(),
	)
}
