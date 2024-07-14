package entity

import (
	"time"
)

type GuestDb struct {
	Uuid            string    `ch:"uuid"`
	GuestHash       string    `ch:"guest_hash"`
	Timestamp       time.Time `ch:"timestamp"`
	Favorites       uint8     `ch:"favorites"`
	Events          uint32    `ch:"events"`
	Sessions        uint32    `ch:"sessions"`
	Hits            uint32    `ch:"hits"`
	Repair          uint8     `ch:"repair"`
	FirstSessionId  string    `ch:"first_session_id"`
	FirstDate       time.Time `ch:"first_date"`
	FirstUrlFrom    string    `ch:"first_url_from"`
	FirstUrlTo      string    `ch:"first_url_to"`
	FirstUrlTo404   bool      `ch:"first_url_to_404"`
	FirstSiteId     string    `ch:"first_site_id"`
	FirstAdvUuid    string    `ch:"first_adv_uuid"`
	FirstReferer1   string    `ch:"first_referer1"`
	FirstReferer2   string    `ch:"first_referer2"`
	FirstReferer3   string    `ch:"first_referer3"`
	LastSessionUuid string    `ch:"last_session_uuid"`
	LastDate        time.Time `ch:"last_date"`
	LastUserId      uint32    `ch:"last_user_id"`
	LastUserAuth    bool      `ch:"last_user_auth"`
	LastUrlLast     string    `ch:"last_url_last"`
	LastUrlLast404  bool      `ch:"last_url_last_404"`
	LastUserAgent   string    `ch:"last_user_agent"`
	LastIp          string    `ch:"last_ip"`
	LastCookie      string    `ch:"last_cookie"`
	LastLanguage    string    `ch:"last_language"`
	LastAdvId       string    `ch:"last_adv_id"`
	LastAdvBack     bool      `ch:"last_adv_back"`
	LastReferer1    string    `ch:"last_referer1"`
	LastReferer2    string    `ch:"last_referer2"`
	LastReferer3    string    `ch:"last_referer3"`
	LastCityId      string    `ch:"last_city_id"`

	//SessionId string    `ch:"session_id"`
	//UrlFrom   string    `ch:"url_from"`
	//UrlTo     string    `ch:"url_to"`
	//UrlTo404  uint8     `ch:"url_to_404"`
	//SiteId    string    `ch:"site_id"`
	//AdvId     string    `ch:"adv_id"`
	//Referer1  string    `ch:"referer1"`
	//Referer2  string    `ch:"referer2"`
	//Referer3  string    `ch:"referer3"`
	//UserId    uint32    `ch:"user_id"`
	//UserAuth  uint8     `ch:"user_auth"`
	//Url       string    `ch:"url"`
	//Url404    uint8     `ch:"url_404"`
	//UserAgent string    `ch:"user_agent"`
	//Ip        string    `ch:"ip"`
	//Cookie    string    `ch:"cookie"`
	//Language  string    `ch:"language"`
	//AdvBack   uint8     `ch:"adv_back"`
	////CountryId  sql.NullString `ch:"country_id"`
	////CityId     sql.NullInt32  `ch:"city_id"`
	////CityInfo   sql.NullString `ch:"city_info"`
}
