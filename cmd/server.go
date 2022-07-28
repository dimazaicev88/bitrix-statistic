package main

import (
	"bitrix-statistic/internal/server"
	"bitrix-statistic/internal/storage"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//TODO 1. Добавить лимит на возврат записей
//TODO 2. Добавить поле CNT_ROWS  в котором будет указано кол-во записей

func main() {
	mysqlStorage := storage.NewMysqlStorage("bitrix", "123", "localhost", "bitrix", 3306)
	serverStatistic := server.NewServer(mysqlStorage)
	err := serverStatistic.Start(3000)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(serverStatistic *server.Server) {
		err := serverStatistic.Stop()
		if err != nil {
			log.Fatalln(err)
		}
	}(serverStatistic)
}
