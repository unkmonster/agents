package data

import (
	"agents/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewSqlxClient)

// Data .
type Data struct {
	// TODO wrapped database client
	db *sqlx.DB
}

func NewSqlxClient(c *conf.Data, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(logger)

	db, err := sqlx.Connect(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
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
