package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DatabaseDSN   string `env:"DATABASE_DSN"`
}

func InitFlags(cfg *Config) {
	flag.StringVar(
		&cfg.ServerAddress, "a", cfg.ServerAddress, "The address where the server will be started",
	)
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "Connection string for database storage")
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	//Инициируем флаги
	InitFlags(cfg)
	flag.Parse()

	//Загружаем переменные окружения
	if err := env.Parse(cfg); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return cfg, nil
}
