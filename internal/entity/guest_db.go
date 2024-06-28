package entity

import (
	"database/sql"
	"time"
)

type GuestDb struct {
	Id             int            `db:"id"`
	CookieToken    string         `db:"cookie_token"`
	TimestampX     sql.NullTime   `db:"timestamp_x"`
	Favorites      string         `db:"favorites"`
	CEvents        int            `db:"c_events"`
	Sessions       int            `db:"sessions"`
	Hits           int            `db:"hits"`
	Repair         string         `db:"repair"`
	FirstSessionId int            `db:"first_session_id"`
	FirstDate      time.Time      `db:"first_date"`
	FirstUrlFrom   sql.NullString `db:"first_url_from"`
	FirstUrlTo     sql.NullString `db:"first_url_to"`
	FirstUrlTo404  sql.NullString `db:"first_url_to_404"`
	FirstSiteId    sql.NullString `db:"first_site_id"`
	FirstAdvId     int            `db:"first_adv_id"`
	FirstReferer1  sql.NullString `db:"first_referer_1"`
	FirstReferer2  sql.NullString `db:"first_referer_2"`
	FirstReferer3  sql.NullString `db:"first_referer_3"`
	LastSessionId  int            `db:"last_session_id"`
	LastDate       time.Time      `db:"last_date"`
	LastUserId     int            `db:"last_user_id"`
	LastUserAuth   sql.NullString `db:"last_user_auth"`
	LastUrlLast    sql.NullString `db:"last_url_last"`
	LastUrlLast404 string         `db:"last_url_last_404"`
	LastUserAgent  sql.NullString `db:"last_user_agent"`
	LastIp         sql.NullString `db:"last_ip"`
	LastCookie     sql.NullString `db:"last_cookie"`
	LastLanguage   sql.NullString `db:"last_language"`
	LastAdvId      sql.NullInt32  `db:"last_adv_id"`
	LastAdvBack    string         `db:"last_adv_back"`
	LastReferer1   sql.NullString `db:"last_referer_1"`
	LastReferer2   sql.NullString `db:"last_referer_2"`
	LastReferer3   sql.NullString `db:"last_referer_3"`
	LastSiteId     sql.NullString `db:"last_site_id"`
	LastCountryId  sql.NullString `db:"last_country_id"`
	LastCityId     sql.NullInt32  `db:"last_city_id"`
	LastCityInfo   sql.NullString `db:"last_city_info"`
}
