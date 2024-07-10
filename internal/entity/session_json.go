package entity

import "time"

type SessionJson struct {
	Id          string    `json:"id,omitempty"`
	GuestId     int       `json:"guest_id,omitempty"`
	Events      int       `json:"c_events,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Hits        int       `json:"hits,omitempty"`
	PhpSessinId string    `json:"phpsessid,omitempty"`
	StopListId  int       `json:"stop_list_id,omitempty"`
}
