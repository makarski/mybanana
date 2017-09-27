package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config holds application configs
type Config struct {
	MySQLDSN string `envconfig:"DB_DSN"`
	Port     int    `envconfig:"PORT" default:"80"`
}

// NewConfig loads application configs
// and returns an Config instance
func NewConfig() (*Config, error) {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, errors.Wrap(err, "NewConfig")
	}
	return &cfg, nil
}
