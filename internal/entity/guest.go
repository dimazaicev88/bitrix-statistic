package entity

import "time"

type Guest struct {
	Id             int       `json:"id,omitempty" db:"id"`
	CookieToken    string    `json:"cookieToken,omitempty" db:"cookie_token"`
	TimestampX     time.Time `json:"timestampX" db:"timestamp_x"`
	Favorites      string    `json:"favorites,omitempty" db:"favorites"`
	CEvents        int       `json:"cEvents,omitempty" db:"c_events"`
	Sessions       int       `json:"sessions,omitempty" db:"sessions"`
	Hits           int       `json:"hits,omitempty" db:"hits"`
	Repair         string    `json:"repair,omitempty" db:"repair"`
	FirstSessionId int       `json:"firstSession_id,omitempty" db:"first_session_id"`
	FirstDate      time.Time `json:"firstDate" db:"first_date"`
	FirstUrlFrom   string    `json:"firstUrl_from,omitempty" db:"first_url_from"`
	FirstUrlTo     string    `json:"firstUrl_to,omitempty" db:"first_url_to"`
	FirstUrlTo404  string    `json:"firstUrl_to_404,omitempty" db:"first_url_to_404"`
	FirstSiteId    string    `json:"firstSite_id,omitempty" db:"first_site_id"`
	FirstAdvId     int       `json:"firstAdv_id,omitempty" db:"first_adv_id"`
	FirstReferer1  string    `json:"firstReferer_1,omitempty" db:"first_referer_1"`
	FirstReferer2  string    `json:"firstReferer_2,omitempty" db:"first_referer_2"`
	FirstReferer3  string    `json:"firstReferer_3,omitempty" db:"first_referer_3"`
	LastSessionId  int       `json:"lastSession_id,omitempty" db:"last_session_id"`
	LastDate       time.Time `json:"lastDate" db:"last_date"`
	LastUserId     int       `json:"lastUser_id,omitempty" db:"last_user_id"`
	LastUserAuth   string    `json:"lastUser_auth,omitempty" db:"last_user_auth"`
	LastUrlLast    string    `json:"lastUrl_last,omitempty" db:"last_url_last"`
	LastUrlLast404 string    `json:"lastUrl_last_404,omitempty" db:"last_url_last_404"`
	LastUserAgent  string    `json:"lastUser_agent,omitempty" db:"last_user_agent"`
	LastIp         string    `json:"lastIp,omitempty" db:"last_ip"`
	LastCookie     string    `json:"lastCookie,omitempty" db:"last_cookie"`
	LastLanguage   string    `json:"lastLanguage,omitempty" db:"last_language"`
	LastAdvId      int       `json:"lastAdv_id,omitempty" db:"last_adv_id"`
	LastAdvBack    string    `json:"lastAdvBack,omitempty" db:"last_adv_back"`
	LastReferer1   string    `json:"lastReferer_1,omitempty" db:"last_referer_1"`
	LastReferer2   string    `json:"lastReferer_2,omitempty" db:"last_referer_2"`
	LastReferer3   string    `json:"lastReferer_3,omitempty" db:"last_referer_3"`
	LastSiteId     string    `json:"lastSite_id,omitempty" db:"last_site_id"`
	LastCountryId  string    `json:"lastCountry_id,omitempty" db:"last_country_id"`
	LastCityId     int       `json:"lastCityId,omitempty" db:"last_city_id"`
	LastCityInfo   string    `json:"lastCityInfo,omitempty" db:"last_city_info"`
}
