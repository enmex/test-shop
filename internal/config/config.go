package config

import "test-shop/pkg/db"

type Config struct {
	DB *db.Config
}

func NewConfig() (*Config, error) {
	dbCfg, err := NewDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB: dbCfg,
	}, nil
}
