package server

import (
	authnv1 "agents/api/authn/service/v1"
	"agents/app/authn/service/internal/conf"
	"agents/app/authn/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, auth *service.AuthnService, middlewares []middleware.Middleware) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middlewares...,
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

	authnv1.RegisterAuthnHTTPServer(srv, auth)
	return srv
}
