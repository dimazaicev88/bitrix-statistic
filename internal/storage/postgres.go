package storage

import "github.com/jmoiron/sqlx"

type PostgresStorage struct {
	DB *sqlx.DB
}

func NewPostgresStorage(login, password, url string) error {

	return nil
}
