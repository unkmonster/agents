package data

import (
	"agents/app/user/service/internal/conf"
	"agents/pkg/migration"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSqlxClient, NewUserDomainRepo, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
}

func NewSqlxClient(c *conf.Data, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(logger)

	db, err := otelsqlx.Open(c.Database.Driver, c.Database.Source,
		otelsql.WithAttributes(semconv.DBSystemMySQL),
	)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	m := migration.New(logger, db.DB, c.Database.Driver, c.Database.MigrationSource)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	return db
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	db := NewSqlxClient(c, logger)

	cleanup := func() {
		log.Info("closing the data resources")
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
