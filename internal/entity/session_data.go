package entity

import "time"

type SessionData struct {
	Id            int       `json:"id,omitempty" db:"id" `
	DateFirst     time.Time `json:"dateFirst,omitempty" db:"date_first"`
	DateLast      time.Time `json:"dateLast,omitempty" db:"date_last"`
	GuestMD5      string    `json:"guestMD5" db:"guest_md5"`
	SessSessionId int       `json:"sessSessionId,omitempty"`
	Token         string    `json:"token,omitempty"`
	SessionDate   string    `json:"sessionData,omitempty" db:"session_data"`
}
