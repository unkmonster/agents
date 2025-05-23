package server

import (
	"agents/app/commission/service/internal/conf"
	"agents/app/commission/service/internal/service"
	"agents/pkg/middleware/basic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	commv1 "agents/api/commission/service/v1"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, comm *service.CommissionService, wallet *service.WalletService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			basic.Server(logger),
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

	srv := http.NewServer(opts...)
	commv1.RegisterCommissionHTTPServer(srv, comm)
	commv1.RegisterWalletHTTPServer(srv, wallet)
	return srv
}
