package entitydb

import "time"

type AdvDay struct {
	Uuid          string    `ch:"uuid"`
	AdvUuid       string    `ch:"adv_uuid"`
	DateStat      time.Time `ch:"date_stat"`
	Guests        uint32    `ch:"guests"`
	GuestsDay     uint32    `ch:"guests_day"`
	NewGuests     uint32    `ch:"new_guests"`
	Favorites     uint32    `ch:"favorites"`
	Hosts         uint32    `ch:"hosts"`
	HostsDay      uint32    `ch:"hosts_day"`
	Sessions      uint32    `ch:"sessions"`
	Hits          uint32    `ch:"hits"`
	GuestsBack    uint32    `ch:"guests_back"`
	GuestsDayBack uint32    `ch:"guests_day_back"`
	FavoritesBack uint32    `ch:"favorites_back"`
	HostsBack     uint32    `ch:"hosts_back"`
	HostsDayBack  uint32    `ch:"hosts_day_back"`
	SessionsBack  uint32    `ch:"sessions_back"`
	HitsBack      uint32    `ch:"hits_back"`
}
