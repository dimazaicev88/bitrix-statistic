package storage

import (
	"bitrix-statistic/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"time"
)

func NewClickHouseClient(cfg config.ServerEnvConfig) *ch.DB {
	db := ch.Connect(
		ch.WithAddr(cfg.ClickHouseHost),
		ch.WithUser(cfg.ClickHouseUser),
		ch.WithPassword(cfg.ClickHousePassword),
		ch.WithDatabase(cfg.ClickHouseDbName),
		ch.WithTimeout(5*time.Second),
		ch.WithDialTimeout(5*time.Second),
		ch.WithReadTimeout(5*time.Second),
		ch.WithWriteTimeout(5*time.Second),
	)

	db.AddQueryHook(chdebug.NewQueryHook(
		chdebug.WithVerbose(true),
		chdebug.FromEnv("CHDEBUG"),
	))
	//
	//conn, err := clickhouse.Open(&clickhouse.Options{
	//	Addr: []string{cfg.ClickHouseHost}, //"127.0.0.1:9000"
	//	Auth: clickhouse.Auth{
	//		Database: cfg.ClickHouseDbName,   //"default"
	//		Username: cfg.ClickHouseUser,     //"default"
	//		Password: cfg.ClickHousePassword, //"dima"
	//	},
	//	Debug: true,
	//	Debugf: func(format string, v ...any) {
	//		logrus.Printf(format+"\n", v...)
	//	},
	//	Settings: clickhouse.Settings{
	//		"max_execution_time": 60,
	//	},
	//	Compression: &clickhouse.Compression{
	//		Method: clickhouse.CompressionLZ4,
	//	},
	//	DialTimeout:          time.Second * 30,
	//	MaxOpenConns:         5,
	//	MaxIdleConns:         5,
	//	ConnMaxLifetime:      time.Duration(10) * time.Minute,
	//	ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
	//	BlockBufferSize:      10,
	//	MaxCompressionBuffer: 10240,
	//	ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
	//		Products: []struct {
	//			Name    string
	//			Version string
	//		}{
	//			{Name: "my-app", Version: "0.1"},
	//		},
	//	},
	//})
	//if err != nil {
	//	return nil, err
	//}

	return db
}
