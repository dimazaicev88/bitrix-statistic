package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func main() {
	users := sq.Select("*").From("users")
	users = users.Where(sq.Or{sq.Eq{"b": 1}, sq.Eq{"b": 2}}).Where(sq.Eq{"c": 3})
	fmt.Println(users.ToSql())
	//sb.AddWhereClause()
	//conn, err := clickhouse.Open(&clickhouse.Options{
	//	Addr: []string{"127.0.0.1:9000"},
	//	Auth: clickhouse.Auth{
	//		Database: "default",
	//		Username: "default",
	//		Password: "dima",
	//	},
	//	Debug: true,
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
	//row := conn.QueryRow(context.Background(), "SELECT uuid FROM guest")
	//var uuid string
	//if err := row.Scan(&uuid); err != nil {
	//	logrus.Fatal(err)
	//}
	//
	////var guestDb entitydb.GuestDb
	////if err := conn.QueryRow(context.Background(), "SELECT * FROM guest").ScanStruct(&guestDb); err != nil {
	////	logrus.Panicln(err)
	////}
	//
	//logrus.Println("result", uuid)
}
