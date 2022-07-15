package builders

import (
	"bitrix-statistic/internal/filters"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountryCase1(t *testing.T) {
	var filter filters.Filter
	jsonStr := `{"SELECT":["SHORT_NAME"],"WHERE":"ID = :countryId","PARAMS":{":countryId":1},"ORDER_BY":["COUNTRY_ID","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	_, sql := NewCountrySQLBuilder(filter).BuildSQL()
	assert.Equal(t, "SELECT t1.SHORT_NAME FROM b_stat_country t1  WHERE  t1.ID =  ?  ORDER BY  t1.COUNTRY_ID , t1.ID ASC", sql.SQL)
	expectedSlice := []interface{}{1, 12}
	if len(sql.Params) != len(expectedSlice) {
		assert.Error(t, errors.New("slice not equal by size"))
	} else {
		for i := range expectedSlice {
			if sql.Params[i] != expectedSlice[i] {
				assert.Error(t, errors.New("value not equal"), sql.Params[i], expectedSlice[i])
			}
		}
	}
}
