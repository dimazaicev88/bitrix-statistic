package builders

import (
	"bitrix-statistic/internal/filters"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterById(t *testing.T) {
	var filter filters.BitrixHitFilter
	jsonStr := `{
       "filter": {
			"ID": 1
		}
	}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, err := NewHitBitrixSQLBuilder(filter).BuildSQL()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "SELECT t_hit.ID,t_hit.SESSION_ID,t_hit.GUEST_ID,t_hit.NEW_GUEST,t_hit.USER_ID,t_hit.USER_AUTH,t_hit.URL,t_hit.URL_404,t_hit.URL_FROM,t_hit.IP,t_hit.METHOD,t_hit.COOKIES, t_hit.USER_AGENT,t_hit.STOP_LIST_ID,t_hit.COUNTRY_ID,t_hit.CITY_ID, t_city.REGION REGION_NAME, t_city.NAME CITY_NAME, t_hit.SITE_ID,DATE_FORMAT(t_hit.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT from b_stat_hit t_hit LEFT JOIN b_stat_city t_city ON (t_city.ID = t_hit.CITY_ID) where  t_hit.ID=1", sql)
}

func TestFilterByUserLike(t *testing.T) {
	var filter filters.BitrixHitFilter
	jsonStr := `{
       "filter": {
			"USER": "125",
            "USER_EXACT_MATCH":"N"
		}
	}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, err := NewHitBitrixSQLBuilder(filter).BuildSQL()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "SELECT t_hit.ID,t_hit.SESSION_ID,t_hit.GUEST_ID,t_hit.NEW_GUEST,t_hit.USER_ID,t_hit.USER_AUTH,t_hit.URL,t_hit.URL_404,t_hit.URL_FROM,t_hit.IP,t_hit.METHOD,t_hit.COOKIES, t_hit.USER_AGENT,t_hit.STOP_LIST_ID,t_hit.COUNTRY_ID,t_hit.CITY_ID, t_city.REGION REGION_NAME, t_city.NAME CITY_NAME, t_hit.SITE_ID,DATE_FORMAT(t_hit.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT from b_stat_hit t_hit LEFT JOIN b_stat_city t_city ON (t_city.ID = t_hit.CITY_ID)LEFT JOIN b_user t_user ON (t_user.ID = t_hit.USER_ID) where  t_hit.USER_ID like \"%125%\" and t_user.LOGIN like \"%125%\" and t_user.LAST_NAME like \"%125%\" and t_user.NAME like \"%125%\" ", sql)
}

func TestFilterByUser(t *testing.T) {
	var filter filters.BitrixHitFilter
	jsonStr := `{
       "filter": {
			"USER": "125"
		}
	}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, err := NewHitBitrixSQLBuilder(filter).BuildSQL()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "SELECT t_hit.ID,t_hit.SESSION_ID,t_hit.GUEST_ID,t_hit.NEW_GUEST,t_hit.USER_ID,t_hit.USER_AUTH,t_hit.URL,t_hit.URL_404,t_hit.URL_FROM,t_hit.IP,t_hit.METHOD,t_hit.COOKIES, t_hit.USER_AGENT,t_hit.STOP_LIST_ID,t_hit.COUNTRY_ID,t_hit.CITY_ID, t_city.REGION REGION_NAME, t_city.NAME CITY_NAME, t_hit.SITE_ID,DATE_FORMAT(t_hit.DATE_HIT, '%d.%m.%Y %H:%i:%s') as  DATE_HIT from b_stat_hit t_hit LEFT JOIN b_stat_city t_city ON (t_city.ID = t_hit.CITY_ID)LEFT JOIN b_user t_user ON (t_user.ID = t_hit.USER_ID) where  t_hit.USER_ID=125 and  t_user.LOGIN=\"125\" and t_user.LAST_NAME=\"125\" and t_user.NAME=\"125\" ", sql)
}
