package entitydb

import "time"

type EventList struct {
	Uuid          string    `ch:"uuid"`
	EventUuid     string    `ch:"event_uuid"`
	Event3        string    `ch:"event3"`
	Money         float64   `ch:"money"`
	DateEnter     time.Time `ch:"date_enter"`
	refererUrl    string    `ch:"referer_url"`
	URL           string    `ch:"url"`
	RedirectUrl   string    `ch:"redirect_url"`
	SessionUuid   string    `ch:"session_uuid"`
	GuestUuid     string    `ch:"guest_uuid"`
	GuestAdvUuid  string    `ch:"guest_adv_uuid"`
	AdvBack       bool      `ch:"adv_back"`
	HitUuid       string    `ch:"hit_uuid"`
	CountryId     string    `ch:"country_id"`
	KeepDays      uint32    `ch:"keep_days"`
	Chargeback    bool      `ch:"chargeback"`
	SiteId        string    `ch:"site_id"`
	RefererSiteId string    `ch:"referer_site_id"`
}
