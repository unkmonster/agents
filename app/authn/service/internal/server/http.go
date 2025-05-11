package server

import (
	authnv1 "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/app/authn/service/internal/service"
	myjwt "agents/pkg/jwt"
	"context"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, auth *service.AuthnService, conf *conf.Auth) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	method, err := myjwt.ConvertSigningMethod(conf.SigningMethod)
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	opts = append(opts, http.Middleware(
		recovery.Recovery(),
		logging.Server(logger),

		selector.Server(jwt.Server(
			func(token *jwtv5.Token) (interface{}, error) {
				return []byte(conf.PublicKey), nil
			}, jwt.WithSigningMethod(method),
		)).Match(func(ctx context.Context, operation string) bool {
			return operation != "/api.authn.service.v1.Authn/Login"
		}).Build(),

		validate.ProtoValidate(),
	))

	srv := http.NewServer(opts...)

	authnv1.RegisterAuthnHTTPServer(srv, auth)
	return srv
}
