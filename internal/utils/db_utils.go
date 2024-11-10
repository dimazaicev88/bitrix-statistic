package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/go-clickhouse/ch"
)

func TruncateTable(ctx context.Context, tableName string, chClient *ch.DB) {
	_, err := chClient.NewTruncateTable().Table(tableName).Exec(ctx)
	if err != nil {
		logrus.Fatal(err)
	}
}
