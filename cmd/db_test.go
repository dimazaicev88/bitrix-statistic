package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

func BenchmarkDb(b *testing.B) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "24zda#1312", "localhost", 3306, "test")
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < b.N; i++ {
		row := db.QueryRow("select id from speed")
		var id int
		err = row.Scan(&id)
	}
}
