package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Guest struct {
	Uuid    uuid.UUID `ch:"uuid"`
	DateAdd time.Time `ch:"date_add"`
	Repair  bool      `ch:"repair"`
}
