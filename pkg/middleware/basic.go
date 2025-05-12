package middleware

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	PublicKey string
	Method    string
}

func convertSigningMethod(method string) (jwtv5.SigningMethod, error) {
	switch strings.ToUpper(method) {
	case "RS256":
		return jwtv5.SigningMethodRS256, nil
	default:
		return nil, fmt.Errorf("unsupported signing method: %s", method)
	}
}

// ServerBasic 返回基础中间件列表
// 1. recovery
// 2. logging
// 3. jwt
// 4. validate
func ServerBasic(logger log.Logger) kratosMiddleware.Middleware {
	return kratosMiddleware.Chain(
		recovery.Recovery(),
		logging.Server(logger),
		validate.ProtoValidate(),
	)
}
