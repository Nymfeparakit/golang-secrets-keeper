package config

import (
	"flag"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DatabaseDSN   string `env:"DATABASE_DSN"`
	EnableHTTPS   bool   `env:"ENABLE_HTTPS" envDefault:"false" json:"enable_https"`
}

func (c *Config) InitFlags() {
	flag.StringVar(
		&c.ServerAddress, "a", c.ServerAddress, "The address where the server will be started",
	)
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Connection string for database storage")
	flag.BoolVar(&c.EnableHTTPS, "s", c.EnableHTTPS, "Should https be used")
}
