package config

import (
	"os"
	"log"
	"errors"

	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v11"
)

type Config struct{
	Port  string `env:"PORT" envDefault:"8080"`
	DBUrl string `env:"DB_URL required notEmpty"`
}

func Load() (*Config) {
	envErr := godotenv.Load()
	if envErr != nil {
		if errors.Is(envErr, os.ErrNotExist) {
			log.Println(".env file not found, reading from system environment variables")
		} else {
			log.Fatalf("Failed to load .env: %v", envErr)
		}
	}
	var cfg Config
	cfgErr := env.Parse(&cfg)
	if cfgErr != nil {
		log.Fatalf("Failed to load config: %v")
	}

	if cfg.Port != "" && cfg.Port[0] != ':' {
		cfg.Port = ":" + cfg.Port
	}

	return &cfg
}
