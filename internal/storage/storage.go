package storage

import "github.com/jmoiron/sqlx"

type Storage interface {
	DB() *sqlx.DB
	Close() error
}
