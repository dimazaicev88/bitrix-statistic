package entity

import "time"

type Guest struct {
	Id             int       `json:"id,omitempty" db:"id"`
	TimestampX     time.Time `json:"timestamp_x" db:"timestamp_x"`
	Favorites      string    `json:"favorites,omitempty" db:"favorites"`
	CEvents        int       `json:"c_events,omitempty" db:"c_events"`
	Sessions       int       `json:"sessions,omitempty" db:"sessions"`
	Hits           int       `json:"hits,omitempty" db:"hits"`
	Repair         string    `json:"repair,omitempty" db:"repair"`
	FirstSessionId int       `json:"first_session_id,omitempty" db:"first_session_id"`
	FirstDate      time.Time `json:"first_date" db:"first_date"`
	FirstUrlFrom   string    `json:"first_url_from,omitempty" db:"first_url_from"`
	FirstUrlTo     string    `json:"first_url_to,omitempty" db:"first_url_to"`
	FirstUrlTo404  string    `json:"first_url_to_404,omitempty" db:"first_url_to_404"`
	FirstSiteId    string    `json:"first_site_id,omitempty" db:"first_site_id"`
	FirstAdvId     int       `json:"first_adv_id,omitempty" db:"first_adv_id"`
	FirstReferer1  string    `json:"first_referer_1,omitempty" db:"first_referer_1"`
	FirstReferer2  string    `json:"first_referer_2,omitempty" db:"first_referer_2"`
	FirstReferer3  string    `json:"first_referer_3,omitempty" db:"first_referer_3"`
	LastSessionId  int       `json:"last_session_id,omitempty" db:"last_session_id"`
	LastDate       time.Time `json:"last_date" db:"last_date"`
	LastUserId     int       `json:"last_user_id,omitempty" db:"last_user_id"`
	LastUserAuth   string    `json:"last_user_auth,omitempty" db:"last_user_auth"`
	LastUrlLast    string    `json:"last_url_last,omitempty" db:"last_url_last"`
	LastUrlLast404 string    `json:"last_url_last_404,omitempty" db:"last_url_last_404"`
	LastUserAgent  string    `json:"last_user_agent,omitempty" db:"last_user_agent"`
	LastIp         string    `json:"last_ip,omitempty" db:"last_ip"`
	LastCookie     string    `json:"last_cookie,omitempty" db:"last_cookie"`
	LastLanguage   string    `json:"last_language,omitempty" db:"last_language"`
	LastAdvId      int       `json:"last_adv_id,omitempty" db:"last_adv_id"`
	LastAdvBack    string    `json:"last_adv_back,omitempty" db:"last_adv_back"`
	LastReferer1   string    `json:"last_referer_1,omitempty" db:"last_referer_1"`
	LastReferer2   string    `json:"last_referer_2,omitempty" db:"last_referer_2"`
	LastReferer3   string    `json:"last_referer_3,omitempty" db:"last_referer_3"`
	LastSiteId     string    `json:"last_site_id,omitempty" db:"last_site_id"`
	LastCountryId  string    `json:"last_country_id,omitempty" db:"last_country_id"`
	LastCityId     int       `json:"last_city_id,omitempty" db:"last_city_id"`
	LastCityInfo   string    `json:"last_city_info,omitempty" db:"last_city_info"`
}
