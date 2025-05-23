package data

import (
	"agents/app/commission/service/internal/conf"
	"agents/pkg/client"
	"agents/pkg/migration"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"github.com/jmoiron/sqlx"

	userv1 "agents/api/user/service/v1"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewUserServiceClient, NewUserRepo, NewCommissionRepo, NewWalletRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
	uc userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userv1.UserClient) (*Data, func(), error) {
	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.NewHelper(logger).Fatalf("failed to connect to db: %v", err)
	}

	m := migration.New(logger, db.DB, c.Database.Driver, c.Database.MigrationSource)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.NewHelper(logger).Fatal(err)
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}
	return &Data{
		db: db,
		uc: uc,
	}, cleanup, nil
}

func NewDiscovery(registry *conf.Registry) registry.Discovery {
	// new consul client

	c := api.DefaultConfig()
	c.Address = registry.Consul.Address
	c.Scheme = registry.Consul.Schema
	client, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)

	return dis
}

func NewUserServiceClient(dis registry.Discovery, logger log.Logger) userv1.UserClient {
	return client.NewUserServiceClient(dis, logger)
}
