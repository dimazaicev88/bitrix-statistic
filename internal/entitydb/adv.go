package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Adv struct {
	Uuid        uuid.UUID `ch:"uuid"`
	Referer1    string    `ch:"referer1"`
	Referer2    string    `ch:"referer2"`
	Cost        float64   `ch:"cost"`
	DateCreated time.Time `ch:"date_create"`
	EventsView  string    `ch:"events_view"`
	Description string    `ch:"description"`
	Priority    uint32    `ch:"priority"`
	//Referer3    string    `ch:"referer3"`
}

type AdvStat struct {
	AdvUuid       uuid.UUID `ch:"adv_uuid"`
	Guests        uint32    `ch:"guests"`
	NewGuests     uint32    `ch:"new_guests"`
	Favorites     uint32    `ch:"favorites"`
	Hosts         uint32    `ch:"hosts"`
	Sessions      uint32    `ch:"sessions"`
	Hits          uint32    `ch:"hits"`
	GuestsBack    uint32    `ch:"guests_back"`
	FavoritesBack uint32    `ch:"favorites_back"`
	HostsBack     uint32    `ch:"hosts_back"`
	SessionsBack  uint32    `ch:"sessions_back"`
	HitsBack      uint32    `ch:"hits_back"`
}
