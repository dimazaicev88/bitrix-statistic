package tasks

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"testing"
	"time"
)

func BenchmarkAddTask(b *testing.B) {

	//// create a cache with capacity equal to 10000 elements
	//cache, err := otter.MustBuilder[string, string](10_000).
	//	CollectStats().
	//	Cost(func(key string, value string) uint32 {
	//		return 1
	//	}).
	//	WithTTL(time.Hour).
	//	Build()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// set item with ttl (1 hour)
	//cache.Set("key", "value")
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	cache.Set("my-unique-key", []byte("value"))
	//fmt.Println(string(entry))
	for i := 0; i < b.N; i++ {
		//cache.Get("key")
		cache.Get("my-unique-key")
	}
	//cache.Close()

	//conn, err := clickhouse.Open(&clickhouse.Options{
	//	Addr: []string{"127.0.0.1:9000"},
	//	Auth: clickhouse.Auth{
	//		Database: "default",
	//		Username: "default",
	//		Password: "dima",
	//	},
	//	Debug: false,
	//	Debugf: func(format string, v ...any) {
	//		fmt.Printf(format+"\n", v...)
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
	//	logrus.Fatal(err)
	//}
	//
	//for i := 0; i < b.N; i++ {
	//	var uuid string
	//	//if err := conn.QueryRow(context.Background(), "SELECT * FROM guest").ScanStruct(&guestDb); err != nil {
	//	//	logrus.Panicln(err)
	//	//}
	//
	//	row := conn.QueryRow(context.Background(), "SELECT uuid FROM guest")
	//	//var
	//	if err := row.Scan(&uuid); err != nil {
	//		logrus.Fatal(err)
	//	}
	//
	//}
}
