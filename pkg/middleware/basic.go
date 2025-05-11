package middleware

import (
	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	PublicKey string
	Method    jwtv5.SigningMethod
}

// ServerBasic 返回基础中间件列表
// 1. recovery
// 2. logging
// 3. jwt
// 4. validate
func ServerBasic(logger log.Logger, jwtConf *JwtConfig) kratosMiddleware.Middleware {
	return kratosMiddleware.Chain(
		recovery.Recovery(),
		logging.Server(logger),
		jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return []byte(jwtConf.PublicKey), nil
		}, jwt.WithSigningMethod(jwtConf.Method)),
		validate.ProtoValidate(),
	)
}
