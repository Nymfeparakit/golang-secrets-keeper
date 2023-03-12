package config

import (
	"flag"
)

type Config struct {
	ServerAddress string `env:"REMOTE_SERVER_ADDRESS" envDefault:":8080"`
}

func (c *Config) InitFlags() {
	flag.StringVar(
		&c.ServerAddress, "a", c.ServerAddress, "The address where the server will be started",
	)
}
