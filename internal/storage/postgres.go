package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type MysqlStorage struct {
	db *sqlx.DB
}

func NewMysqlStorage(login, password, host, dbname string, port int) *MysqlStorage {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", login, password, host, port, dbname)
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
