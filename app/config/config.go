package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct{
	Port  int    `env:"PORT,required,notEmpty"`
	DBUrl string `env:"DB_URL"`
}

func Load() (*Config) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
