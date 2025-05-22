package server

import (
	"agents/app/authn/service/internal/biz"
	"agents/app/authn/service/internal/conf"
	"context"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewJwtKeyFunc)

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

func NewJwtKeyFunc(auth *biz.AuthUserCase) jwtv5.Keyfunc {
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
