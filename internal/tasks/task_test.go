package tasks

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func BenchmarkAddTask(b *testing.B) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "dima",
		},
		Debug: false,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var uuid string
		//if err := conn.QueryRow(context.Background(), "SELECT * FROM guest").ScanStruct(&guestDb); err != nil {
		//	logrus.Panicln(err)
		//}

		row := conn.QueryRow(context.Background(), "SELECT uuid FROM guest")
		//var
		if err := row.Scan(&uuid); err != nil {
			logrus.Fatal(err)
		}

	}
}