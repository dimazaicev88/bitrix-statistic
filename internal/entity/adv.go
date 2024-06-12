package entity

import "time"

type Adv struct {
	Id            int       `json:"id,omitempty" db:"id"`
	Referer1      string    `json:"referer_1,omitempty" db:"referer_1"`
	Referer2      string    `json:"referer_2,omitempty" db:"referer_2"`
	Cost          float64   `json:"cost,omitempty" db:"cost"`
	Revenue       float64   `json:"revenue,omitempty" db:"revenue"`
	EventsView    string    `json:"events_view,omitempty" db:"events_view"`
	Guests        int       `json:"guests,omitempty" db:"guests"`
	NewGuests     int       `json:"new_guests,omitempty" db:"new_guests"`
	Favorites     int       `json:"favorites,omitempty" db:"favorites"`
	CHosts        int       `json:"c_hosts,omitempty" db:"c_hosts"`
	Sessions      int       `json:"sessions,omitempty" db:"sessions"`
	Hits          int       `json:"hits,omitempty" db:"hits"`
	DateFirst     time.Time `json:"date_first" db:"date_first"`
	DateLast      time.Time `json:"date_last" db:"date_last"`
	GuestsBack    int       `json:"guests_back,omitempty" db:"guests_back"`
	FavoritesBack int       `json:"favorites_back,omitempty" db:"favorites_back"`
	HostsBack     int       `json:"hosts_back,omitempty" db:"hosts_back"`
	SessionsBack  int       `json:"sessions_back,omitempty" db:"sessions_back"`
	HitsBack      int       `json:"hits_back,omitempty" db:"hits_back"`
	Description   string    `json:"description,omitempty" db:"description"`
	Priority      int       `json:"priority,omitempty" db:"priority"`
}
