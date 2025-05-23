package server

import (
	"agents/app/stats/service/internal/conf"
	"agents/app/stats/service/internal/service"
	"agents/pkg/middleware/basic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	statsv1 "agents/api/stats/service/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, stats *service.StatsService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			basic.Server(logger),
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

	srv := grpc.NewServer(opts...)
	statsv1.RegisterStatsServer(srv, stats)
	return srv
}
