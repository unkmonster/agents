package data

import (
	"agents/app/order/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
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

	cleanup := func() {
		log.Info("closing the data resources")
		db.Close()
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
