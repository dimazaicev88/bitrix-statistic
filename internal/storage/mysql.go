package storage

import (
	"bitrix-statistic/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type MysqlStorage struct {
	db *sqlx.DB
}

func NewMysqlStorage(cfg config.ServerEnvConfig) *MysqlStorage {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", cfg.StorageUser, cfg.StoragePassword, cfg.StorageHost, cfg.StoragePort, cfg.StorageDbName)
	log.Println("Mysql connections string: " + dataSource)
	dbConn, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Fatalln(err)
	}
	return &MysqlStorage{
		db: dbConn,
	}
	return nil
}

func (ms MysqlStorage) DB() *sqlx.DB {
	return ms.db
}

func (ms MysqlStorage) Close() error {
	err := ms.db.Close()
	if err != nil {
		return err
	}
	return nil
}
