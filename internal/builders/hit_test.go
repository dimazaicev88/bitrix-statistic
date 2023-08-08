package builders

import (
	"bitrix-statistic/internal/filters"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHitCase1(t *testing.T) {
	var filter filters.Filter
	jsonStr := `{"SELECT":["USER","CountryId","CountryId"],"WHERE":"CountryId>:countryId and SessionId>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["CountryId","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, err := NewHitSQLBuilder(filter).BuildSQL()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "SELECT  t1.CountryId , t2.LOGIN, t2.NAME  FROM b_stat_hit t1  INNER JOIN b_stat_country t3 ON (t3.ID = t1.CountryId) LEFT JOIN b_user t2 ON (t2.ID = t1.USER_ID) WHERE  t1.CountryId > ?  and  t1.SessionId > ?  ORDER BY  t1.CountryId , t1.ID ASC", sql.SQL)
	expectedSlice := []interface{}{1, 12}
	if len(sql.Params) != len(expectedSlice) {
		assert.Error(t, errors.New("slice not equal by size"))
	}
	for i := range expectedSlice {
		if sql.Params[i] != expectedSlice[i] {
			assert.Error(t, errors.New("value not equal"), sql.Params[i], expectedSlice[i])
		}
	}
}

func TestHitCase2(t *testing.T) {
	var filter filters.Filter
	jsonStr := `{
  "SELECT": [
    "CountryId"
  ]
}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, _ := NewHitSQLBuilder(filter).BuildSQL()
	assert.Equal(t, "SELECT  t1.CountryId  FROM b_stat_hit t1  INNER JOIN b_stat_country t3 ON (t3.ID = t1.CountryId)", sql.SQL)
	var expectedSlice []interface{}
	if len(sql.Params) != len(expectedSlice) {
		assert.Error(t, errors.New("slice not equal by size"))
	}
}

func TestHitCase3(t *testing.T) {
	var filter filters.Filter
	jsonStr := `{
    "SELECT": []
}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, _ := NewHitSQLBuilder(filter).BuildSQL()
	assert.Equal(t, "SELECT *  FROM b_stat_hit t1 ", sql.SQL)
	var expectedSlice []interface{}
	if len(sql.Params) != len(expectedSlice) {
		assert.Error(t, errors.New("slice not equal by size"))
	}
}

func TestHitCase4(t *testing.T) {
	var filter filters.Filter
	jsonStr := `{}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, _ := NewHitSQLBuilder(filter).BuildSQL()
	assert.Equal(t, "SELECT *  FROM b_stat_hit t1 ", sql.SQL)
	var expectedSlice []interface{}
	if len(sql.Params) != len(expectedSlice) {
		assert.Error(t, errors.New("slice not equal by size"))
	}
}

func BenchmarkBuildHit(b *testing.B) {
	var filter filters.Filter
	jsonStr := `{"SELECT":["CountryId","USER","Country"],"WHERE":"CountryId>:countryId and SessionId>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["CountryId","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	for i := 0; i < b.N; i++ {
		NewHitSQLBuilder(filter).BuildSQL()
	}
	// Здесь выполняем функцию для тестирования
}
