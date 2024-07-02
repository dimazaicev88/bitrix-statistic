package builders

import (
	"bitrix-statistic/internal/filters"
	"bitrix-statistic/internal/utils"
	"errors"
)

type GuestSQLBuilder struct {
	filter filters.Filter
}

var guestFields = map[string]string{
	"id":                "g.id",
	"c_events":          "g.c_events",
	"first_site_id":     "g.first_site_id",
	"last_site_id":      "g.last_site_id",
	"sessions":          "g.sessions",
	"hits":              "g.hits",
	"favorites":         "g.favorites",
	"first_url_from":    "g.first_url_from",
	"first_url_to":      "g.first_url_to",
	"first_url_to_404":  "g.first_url_to_404",
	"first_adv_id":      "g.first_adv_id",
	"first_referer1":    "g.first_referer1",
	"first_referer2":    "g.first_referer2",
	"first_referer3":    "g.first_referer3",
	"last_adv_id":       "g.last_adv_id",
	"last_adv_back":     "g.last_adv_back",
	"last_referer1":     "g.last_referer1",
	"last_referer2":     "g.last_referer2",
	"last_referer3":     "g.last_referer3",
	"last_user_id":      "g.last_user_id",
	"last_user_auth":    "g.last_user_auth",
	"last_url_last":     "g.last_url_last",
	"last_url_last_404": "g.last_url_last_404",
	"last_user_agent":   "g.last_user_agent",
	"last_ip":           "g.last_ip",
	"last_language":     "g.last_language",
	"last_country_id":   "g.last_country_id",
	"last_city_id":      "g.last_city_id",
	"first_date":        "g.first_date",
	"last_date":         "g.last_date",
	"first_session_id":  "g.first_session_id",
	"last_session_id":   "g.last_session_id",
	"region":            "city.region",
	"name":              "city.name",
}

func NewGuestBuilder(filter filters.Filter) GuestSQLBuilder {
	return GuestSQLBuilder{filter: filter}
}

func (g GuestSQLBuilder) Select() (string, error) {
	for _, field := range g.filter.Fields {
		if _, ok := guestFields[field]; !ok {
			return "", errors.New("unknown field: " + field)
		}
	}
	return "", nil
}

func (g GuestSQLBuilder) Where() (string, error) {
	return "", nil
}

func (g GuestSQLBuilder) ToString() (string, error) {
	selectFields, err := g.Select()
	if err != nil {
		return "", err
	}

	where, err := g.Where()
	if err != nil {
		return "", err
	}
	return utils.StringConcat(selectFields, where), nil
}
