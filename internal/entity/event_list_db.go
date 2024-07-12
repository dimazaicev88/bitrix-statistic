package entity

import "time"

type EventListDB struct {
	uuid          string    `ch:"uuid"`
	eventUuid     string    `ch:"event_uuid"`
	event3        string    `ch:"event3"`
	money         float64   `ch:"money"`
	dateEnter     time.Time `ch:"date_enter"`
	refererUrl    string    `ch:"referer_url"`
	url           string    `ch:"url"`
	redirectUrl   string    `ch:"redirect_url"`
	sessionUuid   string    `ch:"session_uuid"`
	guestUuid     string    `ch:"guest_uuid"`
	guestAdvUuid  string    `ch:"guest_adv_uuid"`
	advBack       bool      `ch:"adv_back"`
	hitUuid       string    `ch:"hit_uuid"`
	countryId     string    `ch:"country_id"`
	keepDays      uint32    `ch:"keep_days"`
	chargeback    bool      `ch:"chargeback"`
	siteId        string    `ch:"site_id"`
	refererSiteId string    `ch:"referer_site_id"`
}
