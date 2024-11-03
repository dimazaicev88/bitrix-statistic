package utils

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/sirupsen/logrus"
)

var listTables = []string{
	"guest", "hit", "session",
}

func TruncateAllTables(chClient driver.Conn) {
	for _, table := range listTables {
		err := chClient.Exec(context.Background(), "TRUNCATE "+table)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func TruncateTable(tableName string, chClient driver.Conn) {
	err := chClient.Exec(context.Background(), "TRUNCATE "+tableName)
	if err != nil {
		logrus.Fatal(err)
	}
}
