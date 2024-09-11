package storage

import (
	"bitrix-statistic/internal/config"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"sync"
)

var lock = &sync.Mutex{}

type chStorage struct {
	Client driver.Conn
}

var instance *chStorage

func CHStorage() *chStorage {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		err := godotenv.Load()
		if err != nil {
			logrus.Fatal("Error loading .env file", err.Error())
		}
		chClient, err := NewClickHouseClient(config.GetServerConfig())
		if err != nil {
			logrus.Fatal(err)
		}
		instance = &chStorage{
			Client: chClient,
		}
	}
	return instance
}
