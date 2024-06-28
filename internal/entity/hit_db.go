package entity

import "database/sql"

type HitJson struct {
	Id         int            `json:"db"`
	SessionId  int            `json:"session_id,omitempty"`
	GuestId    sql.NullInt32  `json:"guest_id,omitempty"`
	NewGuest   string         `json:"new_guest"`
	UserId     sql.NullInt32  `json:"user_id,omitempty"`
	UserAuth   sql.NullString `json:"user_auth,omitempty"`
	Url        sql.NullString `json:"url,omitempty"`
	Url404     string         `json:"url404,omitempty"`
	UrlFrom    sql.NullString `json:"url_from,omitempty"`
	Ip         sql.NullString `json:"ip,omitempty"`
	Method     sql.NullString `json:"method,omitempty"`
	Cookies    sql.NullString `json:"cookies,omitempty"`
	UserAgent  sql.NullString `json:"user_agent,omitempty"`
	StopListId sql.NullInt32  `json:"stop_list_id,omitempty"`
	CountryId  sql.NullInt32  `json:"country_id,omitempty"`
	CityId     sql.NullInt32  `json:"city_id,omitempty"`
	RegionName sql.NullString `json:"region_name,omitempty"`
	CityName   sql.NullString `json:"city_name,omitempty"`
	SiteId     sql.NullString `json:"site_id,omitempty"`
	DateHit    sql.NullString `json:"date_hit,omitempty"`
}
