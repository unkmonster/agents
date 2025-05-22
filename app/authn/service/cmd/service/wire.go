//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"agents/app/authn/service/internal/biz"
	"agents/app/authn/service/internal/conf"
	"agents/app/authn/service/internal/data"
	"agents/app/authn/service/internal/server"
	"agents/app/authn/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger, *conf.Auth, *conf.Registry, *conf.Kong, *conf.SystemUser) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
