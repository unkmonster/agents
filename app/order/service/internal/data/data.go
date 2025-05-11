package data

import (
	"agents/app/order/service/internal/conf"
	"agents/pkg/migration"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	}, cleanup, nil
}
