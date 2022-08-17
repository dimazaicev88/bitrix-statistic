package main

import (
	"bitrix-statistic/internal/config"
	"bitrix-statistic/internal/server"
	"bitrix-statistic/internal/storage"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//TODO 1. Добавить лимит на возврат записей
//TODO 2. Добавить поле CNT_ROWS  в котором будет указано кол-во записей

func main() {
	serverConfig := config.ServerEnvConfig{}
	serverConfig.ValidateStorageParams()
	mysqlStorage := storage.NewMysqlStorage(serverConfig)
	serverStatistic := server.NewServer(mysqlStorage)
	err := serverStatistic.Start(125)
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
