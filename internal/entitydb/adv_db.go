package entitydb

import "time"

type AdvDb struct {
	Uuid          string    `ch:"uuid"`
	Referer1      string    `ch:"referer1"`
	Referer2      string    `ch:"referer2"`
	Cost          float64   `ch:"cost"`
	Revenue       float64   `ch:"revenue"`
	EventsView    string    `ch:"events_view"`
	Guests        uint32    `ch:"guests"`
	NewGuests     uint32    `ch:"new_guests"`
	Favorites     uint32    `ch:"favorites"`
	Hosts         uint32    `ch:"hosts"`
	Sessions      uint32    `ch:"sessions"`
	Hits          uint32    `ch:"hits"`
	DateFirst     time.Time `ch:"date_first"`
	DateLast      time.Time `ch:"date_last"`
	GuestsBack    uint32    `ch:"guests_back"`
	FavoritesBack uint32    `ch:"favorites_back"`
	HostsBack     uint32    `ch:"hosts_back"`
	SessionsBack  uint32    `ch:"sessions_back"`
	HitsBack      uint32    `ch:"hits_back"`
	Description   string    `ch:"description"`
	Priority      uint32    `ch:"priority"`
}
