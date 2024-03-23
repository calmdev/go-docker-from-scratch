// package config provides the configuration for the application.
package config

import (
	"github.com/caarlos0/env/v10"
)

// config struct holds the configuration for the application.
type Config struct {
	HttpListenAddress string `env:"HTTP_LISTEN_ADDRESS" envDefault:":3000"`
}

// NewConfig creates a new config struct.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
