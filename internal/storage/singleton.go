package storage

import (
	"bitrix-statistic/internal/config"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/sirupsen/logrus"
	"sync"
)

var chClient driver.Conn
var once sync.Once

func CHClient() driver.Conn {
	once.Do(func() {
		var err error
		_ = NewClickHouseClient(config.GetServerConfig())
		if err != nil {
			logrus.Fatal(err)
		}
	})
	return chClient
}
