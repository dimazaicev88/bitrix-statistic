package models

import (
	"time"
)

type Guest struct {
	GuestHash  string    `ch:"guestHash"`
	DateInsert time.Time `ch:"dateInsert"`
}
