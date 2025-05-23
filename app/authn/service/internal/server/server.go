package server

import (
	"agents/app/authn/service/internal/biz"
	"agents/app/authn/service/internal/conf"
	myjwt "agents/pkg/jwt"
	"context"

	userv1 "agents/api/user/service/v1"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewJwtKeyFunc, NewBeforeStart, NewMiddlewares)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewJwtKeyFunc(auth *biz.AuthUseCase) jwtv5.Keyfunc {
	return func(token *jwtv5.Token) (interface{}, error) {
		sub, err := token.Claims.GetSubject()
		if err != nil {
			return nil, err
		}

		credential, err := auth.GetUserCredential(context.Background(), sub)
		if err != nil {
			return nil, err
		}
		return jwtv5.ParseRSAPublicKeyFromPEM([]byte(*credential.PublicKey))
	}
}

func NewBeforeStart(logger log.Logger, authn *biz.AuthUseCase, conf *conf.SystemUser) func(context.Context) error {
	return func(context.Context) error {
		user, err := authn.RegisterZeroUser(context.Background(), conf.Username, conf.Password)
		if userv1.IsUserIsExists(err) {
			log.NewHelper(logger).Infof("skip register zero user: user was created")
			return nil
		}
		if err != nil {
			log.NewHelper(logger).Warnf("register zero user fail: %v", err)
			return err
		}
		log.NewHelper(logger).Infof("register zero user complete: %#v", user)
		return nil
	}
}

func NewMiddlewares(logger log.Logger, keyfunc jwtv5.Keyfunc) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),

		selector.Server(jwt.Server(
			keyfunc,
			jwt.WithSigningMethod(jwtv5.SigningMethodRS256), // TODO: 签名算法不要硬编码
			jwt.WithClaims(func() jwtv5.Claims { return &myjwt.UserClaims{} }),
		)).Match(func(ctx context.Context, operation string) bool {
			return operation != "/api.authn.service.v1.Authn/Login"
		}).Build(),

		validate.ProtoValidate(),
	}
}
