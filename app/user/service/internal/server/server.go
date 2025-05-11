package server

import (
	"agents/app/user/service/internal/conf"
	"agents/pkg/middleware"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewMiddleware)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	log.Infof("address: %s", conf.Consul.Address)

	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewMiddleware(logger log.Logger, conf *conf.Auth) kratosMiddleware.Middleware {
	return middleware.ServerBasic(logger, &middleware.JwtConfig{
		PublicKey: conf.PublicKey,
		Method:    conf.SigningMethod,
	})
}
