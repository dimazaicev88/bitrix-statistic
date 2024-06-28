package entity

import "time"

type GuestJson struct {
	Id             int       `json:"id,omitempty"`
	CookieToken    string    `json:"cookieToken,omitempty"`
	TimestampX     time.Time `json:"timestampX"`
	Favorites      string    `json:"favorites,omitempty"`
	CEvents        int       `json:"cEvents,omitempty"`
	Sessions       int       `json:"sessions,omitempty"`
	Hits           int       `json:"hits,omitempty"`
	Repair         string    `json:"repair,omitempty"`
	FirstSessionId int       `json:"firstSession_id,omitempty"`
	FirstDate      time.Time `json:"firstDate"`
	FirstUrlFrom   string    `json:"firstUrl_from,omitempty"`
	FirstUrlTo     string    `json:"firstUrl_to,omitempty"`
	FirstUrlTo404  string    `json:"firstUrl_to_404,omitempty"`
	FirstSiteId    string    `json:"firstSite_id,omitempty"`
	FirstAdvId     int       `json:"firstAdv_id,omitempty"`
	FirstReferer1  string    `json:"firstReferer_1,omitempty"`
	FirstReferer2  string    `json:"firstReferer_2,omitempty"`
	FirstReferer3  string    `json:"firstReferer_3,omitempty"`
	LastSessionId  int       `json:"lastSession_id,omitempty"`
	LastDate       time.Time `json:"lastDate"`
	LastUserId     int       `json:"lastUser_id,omitempty"`
	LastUserAuth   string    `json:"lastUser_auth,omitempty"`
	LastUrlLast    string    `json:"lastUrl_last,omitempty"`
	LastUrlLast404 string    `json:"lastUrl_last_404,omitempty"`
	LastUserAgent  string    `json:"lastUser_agent,omitempty"`
	LastIp         string    `json:"lastIp,omitempty"`
	LastCookie     string    `json:"lastCookie,omitempty"`
	LastLanguage   string    `json:"lastLanguage,omitempty"`
	LastAdvId      int       `json:"lastAdv_id,omitempty"`
	LastAdvBack    string    `json:"lastAdvBack,omitempty"`
	LastReferer1   string    `json:"lastReferer_1,omitempty"`
	LastReferer2   string    `json:"lastReferer_2,omitempty"`
	LastReferer3   string    `json:"lastReferer_3,omitempty"`
	LastSiteId     string    `json:"lastSite_id,omitempty"`
	LastCountryId  string    `json:"lastCountry_id,omitempty"`
	LastCityId     int       `json:"lastCityId,omitempty"`
	LastCityInfo   string    `json:"lastCityInfo,omitempty"`
}
