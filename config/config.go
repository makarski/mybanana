package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
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
		return nil, fmt.Errorf("NewConfig: %s", err)
	}
	return &cfg, nil
}
