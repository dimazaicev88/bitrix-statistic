package models

import (
	"github.com/uptrace/go-clickhouse/ch"
	"time"
)

type Guest struct {
	ch.CHModel `ch:"partition:toYYYYMM(dateInsert)"`

	GuestHash  string    `ch:"guestHash"`
	DateInsert time.Time `ch:"dateInsert,pk"`
}
