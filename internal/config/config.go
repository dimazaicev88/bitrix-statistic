package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type ServerEnvConfig struct {
	StorageHost     string `env:"STORAGE_HOST,notEmpty"`
	StorageUser     string `env:"STORAGE_USER,notEmpty"`
	StoragePassword string `env:"STORAGE_PASSWORD,notEmpty"`
	StorageDbName   string `env:"STORAGE_DB_NAME,notEmpty"`
	StoragePort     string `env:"STORAGE_PORT,notEmpty"`
}

func ParseServerConfig() ServerEnvConfig {
	config := ServerEnvConfig{}
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
