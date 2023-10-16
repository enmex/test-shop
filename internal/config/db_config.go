package config

import (
	"os"
	"strconv"
	"test-shop/pkg/db"
)

func NewDBConfig() (*db.Config, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &db.Config{
		Host:               os.Getenv("DB_HOST"),
		Port:               dbPort,
		DatabaseName:       os.Getenv("DB_NAME"),
		User:               os.Getenv("DB_USER"),
		Password:           os.Getenv("DB_PASS"),
		MigrationDirectory: os.Getenv("MIGRATION_DIR"),
		SslMode:            os.Getenv("DB_SSL"),
		DriverName:         os.Getenv("DB_DRIVER"),
	}, nil
}
