package entity

import "time"

type AdvGuestDB struct {
	Uuid         string    `ch:"uuid"`
	AdvUuid      string    `ch:"adv_uuid"`
	Back         bool      `ch:"back"`
	GuestUuid    string    `ch:"guest_uuid"`
	DateGuestHit time.Time `ch:"date_guest_hit"`
	DateHostHit  time.Time `ch:"date_host_hit"`
	SessionUuid  string    `ch:"session_uuid"`
	Ip           string    `ch:"ip"`
}
