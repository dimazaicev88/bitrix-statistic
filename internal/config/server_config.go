package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type ServerEnvConfig struct {
	ClickHouseHost     string `env:"CLICKHOUSE_HOST,notEmpty"`
	ClickHouseUser     string `env:"CLICKHOUSE_USER,notEmpty"`
	ClickHousePassword string `env:"CLICKHOUSE_PASSWORD,notEmpty"`
	ClickHouseDbName   string `env:"CLICKHOUSE_DB_NAME,notEmpty"`
	RedisHost          string `env:"REDIS_HOST,notEmpty"`
	ServerPort         int    `env:"SERVER_PORT,notEmpty"`
	TaskMonitorPort    int    `env:"TASK_MONITOR_PORT,notEmpty"`
}

func GetServerConfig() ServerEnvConfig {
	config := ServerEnvConfig{}
	if err := env.Parse(&config); err != nil {
		logrus.Fatalln(err)
	}
	return config
}
