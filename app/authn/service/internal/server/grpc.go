package server

import (
	authnv1 "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/app/authn/service/internal/service"
	"context"

	myjwt "agents/pkg/jwt"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, authn *service.AuthnService, logger log.Logger, conf *conf.Auth) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	method, err := myjwt.ConvertSigningMethod(conf.SigningMethod)
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	opts = append(opts, grpc.Middleware(
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

	srv := grpc.NewServer(opts...)

	authnv1.RegisterAuthnServer(srv, authn)

	return srv
}
