package consul

import (
	consulv2 "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

type ApiConfig struct {
	Address string
	Schema  string
}

func NewRegistrar(logger log.Logger, cfg *consulAPI.Config) registry.Registrar {
	cli, err := consulAPI.NewClient(cfg)
	if err != nil {
		log.NewHelper(logger).Fatalf("failed to new consul client: %v", err)
	}
	return consulv2.New(cli)
}
