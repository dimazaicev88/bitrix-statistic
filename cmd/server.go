package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/huandu/go-sqlbuilder"
)

//TODO 1. Добавить лимит на возврат записей
//TODO 2. Добавить поле CNT_ROWS  в котором будет указано кол-во записей

func main() {

	sb := sqlbuilder.Select("id", "name").From("demo.user")
	sb.Where("a=12").Where("b=12")
	fmt.Println(sb.String())
	//serverConfig := config.ParseServerConfig()
	//mysqlStorage := storage.NewMysqlStorage(serverConfig)
	//serverStatistic := server.NewServer(mysqlStorage)
	//err := serverStatistic.Start(125)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer func(serverStatistic *server.Server) {
	//	err := serverStatistic.Stop()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}(serverStatistic)
}
