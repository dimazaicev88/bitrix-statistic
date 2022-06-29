package app

import "github.com/jmoiron/sqlx"

type Application struct {
	Storage *sqlx.DB
}

func NewApp() {

}
