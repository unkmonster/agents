package data

import (
	"agents/app/authn/service/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	commv1 "agents/api/commission/service/v1"
	userv1 "agents/api/user/service/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewSqlxClient,
	NewDiscovery,
	NewUserServiceClient,
	NewUserCredentialRepo,
	NewUserRepo,
	NewCommissionServiceClient,
	NewCommissionRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
	uc userv1.UserClient
	cc commv1.CommissionClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *sqlx.DB, uc userv1.UserClient, cc commv1.CommissionClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{
		db: db,
		uc: uc,
		cc: cc,
	}, cleanup, nil
}

func NewSqlxClient(c *conf.Data) *sqlx.DB {
	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return db
}

func NewDiscovery(registry *conf.Registry) registry.Discovery {
	// new consul client
	log.Infof("registry: %v", registry)
	log.Infof("consul: %v", registry.Consul)

	c := api.DefaultConfig()
	c.Address = registry.Consul.Address
	c.Scheme = registry.Consul.Scheme
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)

	return dis
}

func NewUserServiceClient(dis registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///agents.user.service"),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userv1.NewUserClient(conn)
}

func NewCommissionServiceClient(dis registry.Discovery) commv1.CommissionClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///agents.commission.service"),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return commv1.NewCommissionClient(conn)
}
