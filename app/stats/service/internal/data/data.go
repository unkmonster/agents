package data

import (
	"agents/app/stats/service/internal/conf"
	"agents/pkg/client"
	"agents/pkg/migration"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	commv1 "agents/api/commission/service/v1"
	userv1 "agents/api/user/service/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewStatsRepo,
	NewDiscovery,
	NewCommissionRepo,
	NewDomainRepo,
)

// Data .
type Data struct {
	db *sqlx.DB
	uc userv1.UserClient
	cc commv1.CommissionClient
}

func NewDiscovery(logger log.Logger, registry *conf.Registry) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = registry.Consul.Address
	c.Scheme = registry.Consul.Schema
	client, err := api.NewClient(c)
	if err != nil {
		log.NewHelper(logger).Fatalf("new discovery failed: %v", err)
	}
	// new dis with consul client
	return consul.New(client)
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, dis registry.Discovery) (*Data, func(), error) {
	log := log.NewHelper(logger)
	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	m := migration.New(logger, db.DB, c.Database.Driver, c.Database.MigrationSource)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	cleanup := func() {
		log.Info("closing the data resources")
		db.Close()
	}
	return &Data{
		db: db,
		uc: client.NewUserServiceClient(dis, logger),
		cc: client.NewCommissionServiceClient(dis, logger),
	}, cleanup, nil
}
