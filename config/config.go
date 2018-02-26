package config

import (
	"github.com/codingconcepts/env"
)

// Config is a configuration struct populated by environmental variables
type Config struct {
	Address          string `env:"ADDRESS" default:"0.0.0.0"`
	Port             string `env:"PORT" default:"8443"`
	TLSCertPath      string `env:"CERT_PATH" default:"server.crt"`
	TLSKeyPath       string `env:"KEY_PATH" default:"server.key"`
	StatsdEndpoint   string `env:"STATSD_ENDPOINT" required:"false"`
	DatabaseUser     string `env:"DB_USER" default:"postgres"`
	DatabasePassword string `env:"DB_PASSWORD" default:"postgres"`
	DatabaseHost     string `env:"DB_HOST" default:"localhost"`
}

func Load() (config *Config, err error) {
	config = &Config{}

	if err = env.Set(config); err != nil {
		return nil, err
	}

	return
}
