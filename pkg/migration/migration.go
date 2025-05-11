package migration

import (
	"database/sql"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func New(logger log.Logger, db *sql.DB, driver, source string) *migrate.Migrate {
	var inst database.Driver
	var err error

	if strings.ToLower(driver) == "mysql" {
		inst, err = mysql.WithInstance(db, &mysql.Config{})
	} else {
		log.NewHelper(logger).Fatalf("failed to new migrate unsupported driver: %s", driver)
	}

	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		source,
		driver,
		inst,
	)
	if err != nil {
		log.NewHelper(logger).Fatal(err)
	}

	return m
}
