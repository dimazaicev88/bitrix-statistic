package entitydb

import (
	"time"
)

type Guest struct {
	GuestHash string    `ch:"guestHash"`
	DateAdd   time.Time `ch:"dateAdd"`
}
