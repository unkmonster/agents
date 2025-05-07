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

	userv1 "agents/api/user/service/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
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
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
	uc userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *sqlx.DB, uc userv1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{
		db: db,
		uc: uc,
	}, cleanup, nil
}

func NewSqlxClient(c *conf.Data) *sqlx.DB {
	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		panic(err)
	}
	return db
}

func NewDiscovery() registry.Discovery {
	// new consul client
	client, err := api.NewClient(api.DefaultConfig())
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
