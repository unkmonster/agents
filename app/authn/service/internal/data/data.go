package data

import (
	"agents/app/authn/service/internal/conf"
	"agents/pkg/client"
	"agents/pkg/migration"

	"github.com/dubonzi/otelresty"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"

	userv1 "agents/api/user/service/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/consul/api"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewSqlxClient,
	NewDiscovery,
	client.NewUserServiceClient,
	NewUserCredentialRepo,
	NewUserRepo,
	NewGatewayRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *sqlx.DB
	uc  userv1.UserClient
	cli *resty.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *sqlx.DB, uc userv1.UserClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		db.Close()
	}

	return &Data{
		db:  db,
		uc:  uc,
		cli: NewRestyClient(),
	}, cleanup, nil
}

func NewSqlxClient(c *conf.Data, logger log.Logger) *sqlx.DB {
	db, err := otelsqlx.Open(
		c.Database.Driver,
		c.Database.Source,
		//otelsql.WithDBSystem("mysql"),
		//otelsql.WithTracerProvider(otel.GetTracerProvider()),
	)
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	m := migration.New(logger, db.DB, c.Database.Driver, c.Database.MigrationSource)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.NewHelper(logger).Fatal(err)
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

func NewRestyClient() *resty.Client {

	cli := resty.New()
	cli.SetRetryCount(5)

	otelresty.TraceClient(cli)
	return cli
	// cli.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
	// 	// Now you have access to the Client and Request instance
	// 	// manipulate it as per your need

	// 	return nil // if its successful otherwise return error
	// })
}
