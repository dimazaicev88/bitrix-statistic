package utils

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/sirupsen/logrus"
)

func TruncateTable(ctx context.Context, tableName string, chClient driver.Conn) {
	err := chClient.Exec(ctx, "TRUNCATE "+tableName)
	if err != nil {
		logrus.Fatal(err)
	}
}
