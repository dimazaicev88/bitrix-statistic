package entitydb

import (
	"github.com/google/uuid"
	"time"
)

type Guest struct {
	Uuid             uuid.UUID `ch:"uuid"`
	DateAdd          time.Time `ch:"date_add"`
	Favorites        uint8     `ch:"favorites"`
	Events           uint32    `ch:"events"`
	Sessions         uint32    `ch:"sessions"`
	Hits             uint32    `ch:"hits"`
	Repair           uint8     `ch:"repair"`
	PhpSessionId     string    `ch:"php_session_id"`
	FirstSessionUuid string    `ch:"session_uuid"`
	FirstDate        time.Time `ch:"first_date"`
	FirstUrlFrom     string    `ch:"first_url_from"`
	FirstUrlTo       string    `ch:"first_url_to"`
	FirstUrlTo404    bool      `ch:"first_url_to_404"`
	FirstSiteId      string    `ch:"first_site_id"`
	FirstAdvUuid     uuid.UUID `ch:"first_adv_uuid"`
	FirstReferer1    string    `ch:"first_referer1"`
	FirstReferer2    string    `ch:"first_referer2"`
	FirstReferer3    string    `ch:"first_referer3"`
	LastSessionUuid  string    `ch:"last_session_uuid"`
	LastDate         time.Time `ch:"last_last_date"`
	LastUserId       uint32    `ch:"last_user_id"`
	LastUserAuth     bool      `ch:"last_user_auth"`
	LastUrlLast      string    `ch:"last_url_last"`
	LastUrlLast404   string    `ch:"last_url_last_404"`
	LastUserAgent    string    `ch:"last_user_agent"`
	LastIp           string    `ch:"last_ip"`
	LastCookie       string    `ch:"last_cookie"`
	LastLanguage     string    `ch:"last_language"`
	LastAdvUUid      uuid.UUID `ch:"last_adv_uuid"`
	LastAdvBack      bool      `ch:"last_adv_back"`
	LastReferer1     string    `ch:"last_referer1"`
	LastReferer2     string    `ch:"last_referer2"`
	LastReferer3     string    `ch:"last_referer3"`
	LastSiteId       string    `ch:"last_site_id"`
	LastCountryId    string    `ch:"last_country_id"`
	LastCityId       string    `ch:"last_city_id"`
	Sign             int8      `ch:"sign"`
	Version          uint32    `ch:"version"`
	LastCityInfo     uint32    `ch:"last_city_info"`
	GuestHash        string    `ch:"guest_hash"`
}
