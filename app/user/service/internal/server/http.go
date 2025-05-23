package server

import (
	userv1 "agents/api/user/service/v1"
	"agents/app/user/service/internal/conf"
	"agents/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, user *service.UserService, middlewares []middleware.Middleware) *http.Server {
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

	userv1.RegisterUserHTTPServer(srv, user)
	return srv
}
