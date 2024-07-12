package entity

import (
	"time"
)

type GuestDb struct {
	Uuid           string    `ch:"uuid"`
	Token          string    `ch:"token"`
	Timestamp      time.Time `ch:"timestamp"`
	Favorites      uint8     `ch:"favorites"`
	Events         uint32    `ch:"events"`
	Sessions       uint32    `ch:"sessions"`
	Hits           uint32    `ch:"hits"`
	Repair         uint8     `ch:"repair"`
	firstSessionId string    `ch:"first_session_id"`
	firstDate      time.Time `ch:"first_date"`
	firstUrlFrom   string    `ch:"first_url_from"`
	firstUrlTo     string    `ch:"first_url_to"`
	firstUrlTo404  bool      `ch:"first_url_to_404"`
	firstSiteId    string    `ch:"first_site_id"`
	firstAdvId     string    `ch:"first_adv_id"`
	firstReferer1  string    `ch:"first_referer1"`
	firstReferer2  string    `ch:"first_referer2"`
	firstReferer3  string    `ch:"first_referer3"`
	lastSessionId  string    `ch:"last_session_id"`
	lastDate       time.Time `ch:"last_date"`
	lastUserId     uint32    `ch:"last_user_id"`
	lastUserAuth   bool      `ch:"last_user_auth"`
	lastUrlLast    string    `ch:"last_url_last"`
	lastUrlLast404 bool      `ch:"last_url_last_404"`
	lastUserAgent  string    `ch:"last_user_agent"`
	lastIp         string    `ch:"last_ip"`
	lastCookie     string    `ch:"last_cookie"`
	lastLanguage   string    `ch:"last_language"`
	lastAdvId      string    `ch:"last_adv_id"`
	lastAdvBack    bool      `ch:"last_adv_back"`
	lastReferer1   string    `ch:"last_referer1"`
	lastReferer2   string    `ch:"last_referer2"`
	lastReferer3   string    `ch:"last_referer3"`
	lastCityId     string    `ch:"last_city_id"`

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
