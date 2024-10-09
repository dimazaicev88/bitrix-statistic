package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Uuid         uuid.UUID `ch:"uuid"`
	GuestUuid    uuid.UUID `ch:"guest_uuid"`
	PhpSessionId string    `ch:"php_session_id"`
	DateAdd      time.Time `ch:"date_add"`
}
