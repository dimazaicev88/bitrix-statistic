package entity

import (
	"database/sql"
	"time"
)

type GuestDb struct {
	Id          int            `db:"id"`
	CookieToken string         `db:"cookie_token"`
	TimestampX  sql.NullTime   `db:"timestamp_x"`
	Favorites   string         `db:"favorites"`
	Events      int            `db:"events"`
	Sessions    int            `db:"sessions"`
	Hits        int            `db:"hits"`
	Repair      string         `db:"repair"`
	SessionId   int            `db:"session_id"`
	Date        time.Time      `db:"date"`
	UrlFrom     sql.NullString `db:"url_from"`
	UrlTo       sql.NullString `db:"url_to"`
	UrlTo404    sql.NullString `db:"url_to_404"`
	SiteId      sql.NullString `db:"site_id"`
	AdvId       int            `db:"adv_id"`
	Referer1    sql.NullString `db:"referer_1"`
	Referer2    sql.NullString `db:"referer_2"`
	Referer3    sql.NullString `db:"referer_3"`
	UserId      int            `db:"user_id"`
	UserAuth    sql.NullString `db:"user_auth"`
	Url         sql.NullString `db:"url"`
	Url404      string         `db:"url_404"`
	UserAgent   sql.NullString `db:"user_agent"`
	Ip          sql.NullString `db:"ip"`
	Cookie      sql.NullString `db:"cookie"`
	Language    sql.NullString `db:"language"`
	AdvBack     string         `db:"adv_back"`
	CountryId   sql.NullString `db:"country_id"`
	CityId      sql.NullInt32  `db:"city_id"`
	CityInfo    sql.NullString `db:"city_info"`
}
