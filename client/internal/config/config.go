package config

import (
	"flag"
)

// Config - settings for client application.
type Config struct {
	ServerAddress string `env:"REMOTE_SERVER_ADDRESS" envDefault:":8080"`
	EnableHTTPS   bool   `env:"ENABLE_HTTPS" envDefault:"false" json:"enable_https"`
}

// InitFlags creates command line flags for settings.
func (c *Config) InitFlags() {
	flag.StringVar(
		&c.ServerAddress, "a", c.ServerAddress, "The address where the server will be started",
	)
	flag.BoolVar(&c.EnableHTTPS, "s", c.EnableHTTPS, "Should https be used")
}
