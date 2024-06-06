package entity

import "time"

type StatSessionData struct {
	ID            int
	DateFirst     time.Duration
	DateLast      time.Duration
	GuestMd5      string
	SessSessionId int
	SessionData   string
}
