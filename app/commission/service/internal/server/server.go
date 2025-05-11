package server

import (
	"agents/app/commission/service/internal/conf"
	"agents/pkg/consul"
	"agents/pkg/middleware"

	"github.com/go-kratos/kratos/v2/log"
	kratosMiddleware "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewBasicMiddleware)

func NewRegistrar(logger log.Logger, conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Schema

	return consul.NewRegistrar(logger, c)
}

func NewBasicMiddleware(logger log.Logger, auth *conf.Auth) kratosMiddleware.Middleware {
	return middleware.ServerBasic(logger, &middleware.JwtConfig{
		PublicKey: auth.PublicKey,
		Method:    auth.SigningMethod,
	})
}
