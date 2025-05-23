package server

import (
	"agents/app/commission/service/internal/conf"
	"agents/app/commission/service/internal/service"

	commv1 "agents/api/commission/service/v1"

	"agents/pkg/middleware/basic"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger, comm *service.CommissionService, wallet *service.WalletService) *grpc.Server {
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
	commv1.RegisterCommissionServer(srv, comm)
	commv1.RegisterWalletServer(srv, wallet)
	return srv
}
