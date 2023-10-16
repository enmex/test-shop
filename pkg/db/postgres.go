package db

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func Migrations(conf *Config, log *logrus.Logger) error {
	db, err := sql.Open(conf.DriverName, conf.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}
