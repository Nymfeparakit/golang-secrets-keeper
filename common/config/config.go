package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
)

type Config interface {
	InitFlags()
}

func InitConfig(cfg Config) error {
	cfg.InitFlags()
	flag.Parse()

	//Загружаем переменные окружения
	if err := env.Parse(cfg); err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
