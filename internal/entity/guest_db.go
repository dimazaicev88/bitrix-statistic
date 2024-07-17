package entity

import (
	"time"
)

type GuestDb struct {
	Uuid        string    `ch:"uuid"`
	GuestHash   string    `ch:"guest_hash"`
	Timestamp   time.Time `ch:"timestamp"`
	Favorites   uint8     `ch:"favorites"`
	Events      uint32    `ch:"events"`
	Sessions    uint32    `ch:"sessions"`
	Hits        uint32    `ch:"hits"`
	Repair      uint8     `ch:"repair"`
	SessionId   string    `ch:"session_id"`
	Date        time.Time `ch:"date"`
	UrlFrom     string    `ch:"url_from"`
	UrlTo       string    `ch:"url_to"`
	UrlTo404    bool      `ch:"url_to_404"`
	SiteId      string    `ch:"site_id"`
	AdvUuid     string    `ch:"adv_uuid"`
	Referer1    string    `ch:"referer1"`
	Referer2    string    `ch:"referer2"`
	Referer3    string    `ch:"referer3"`
	SessionUuid string    `ch:"session_uuid"`
	UserId      uint32    `ch:"user_id"`
	UserAuth    bool      `ch:"user_auth"`
	UrlLast     string    `ch:"url_last"`
	UserAgent   string    `ch:"user_agent"`
	Ip          string    `ch:"ip"`
	Cookie      string    `ch:"cookie"`
	Language    string    `ch:"language"`
	AdvId       string    `ch:"adv_id"`
	AdvBack     bool      `ch:"adv_back"`
	CityId      string    `ch:"city_id"`
}
