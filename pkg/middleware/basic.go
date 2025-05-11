package middleware

import (
	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
)

// ServerBasic 返回基础中间件列表
// 1. logging
// 2. validate
func ServerBasic(logger log.Logger) kratosMiddleware.Middleware {
	return kratosMiddleware.Chain(
		logging.Server(logger),
		validate.ProtoValidate(),
	)
}
