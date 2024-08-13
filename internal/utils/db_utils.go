package utils

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/sirupsen/logrus"
)

var listTables = []string{
	"adv", "adv_day", "adv_event", "adv_event_day", "adv_guest", "adv_page", "adv_searcher",
	"city", "city_day", "city_ip", "country", "country_day", "day", "event", "event_day", "event_list",
	"guest", "hit", "options", "page", "path", "path_adv", "path_cache", "phrase_list", "searcher_hit",
	"searcher_params", "session",
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
