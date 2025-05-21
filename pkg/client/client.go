package client

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	commv1 "agents/api/commission/service/v1"
	userv1 "agents/api/user/service/v1"
)

func NewUserServiceClient(dis registry.Discovery, logger log.Logger) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///agents.user.service"),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		log.NewHelper(logger).Fatalf("dial \"User\" grpc service failed: %v", err)
	}
	return userv1.NewUserClient(conn)
}

func NewCommissionServiceClient(dis registry.Discovery, logger log.Logger) commv1.CommissionClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///agents.commission.service"),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
		),
	)
	if err != nil {
		log.NewHelper(logger).Fatalf("dial Commission grpc service failed: %v", err)
	}
	return commv1.NewCommissionClient(conn)
}
