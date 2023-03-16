package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
)

// Config - storage for general settings.
type Config interface {
	InitFlags()
}

// InitConfig fills config with data from specified flags and environment variables.
func InitConfig(cfg Config) error {
	cfg.InitFlags()
	flag.Parse()

	if err := env.Parse(cfg); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
