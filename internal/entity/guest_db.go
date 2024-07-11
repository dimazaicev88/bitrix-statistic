package entity

import (
	"time"
)

type GuestDb struct {
	Uuid      string    `ch:"uuid"`
	Token     string    `ch:"token"`
	Timestamp time.Time `ch:"timestamp"`
	Favorites uint8     `ch:"favorites"`
	Events    int32     `ch:"events"`
	Sessions  int32     `ch:"sessions"`
	Hits      int32     `ch:"hits"`
	Repair    uint8     `ch:"repair"`
	SessionId string    `ch:"session_id"`
	UrlFrom   string    `ch:"url_from"`
	UrlTo     string    `ch:"url_to"`
	UrlTo404  uint8     `ch:"url_to_404"`
	SiteId    string    `ch:"site_id"`
	AdvId     string    `ch:"adv_id"`
	Referer1  string    `ch:"referer1"`
	Referer2  string    `ch:"referer2"`
	Referer3  string    `ch:"referer3"`
	UserId    int32     `ch:"user_id"`
	UserAuth  uint8     `ch:"user_auth"`
	Url       string    `ch:"url"`
	Url404    uint8     `ch:"url_404"`
	UserAgent string    `ch:"user_agent"`
	Ip        string    `ch:"ip"`
	Cookie    string    `ch:"cookie"`
	Language  string    `ch:"language"`
	AdvBack   uint8     `ch:"adv_back"`
	//CountryId  sql.NullString `ch:"country_id"`
	//CityId     sql.NullInt32  `ch:"city_id"`
	//CityInfo   sql.NullString `ch:"city_info"`
}
