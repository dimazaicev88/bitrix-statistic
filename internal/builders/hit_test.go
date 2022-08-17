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
	jsonStr := `{"SELECT":["USER","COUNTRY_ID","COUNTRY_ID"],"WHERE":"COUNTRY_ID>:countryId and SESSION_ID>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["COUNTRY_ID","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, err := NewHitSQLBuilder(filter).BuildSQL()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "SELECT  t1.COUNTRY_ID , t2.LOGIN, t2.NAME  FROM b_stat_hit t1  INNER JOIN b_stat_country t3 ON (t3.ID = t1.COUNTRY_ID) LEFT JOIN b_user t2 ON (t2.ID = t1.USER_ID) WHERE  t1.COUNTRY_ID > ?  and  t1.SESSION_ID > ?  ORDER BY  t1.COUNTRY_ID , t1.ID ASC", sql.SQL)
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
    "COUNTRY_ID"
  ]
}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	sql, _ := NewHitSQLBuilder(filter).BuildSQL()
	assert.Equal(t, "SELECT  t1.COUNTRY_ID  FROM b_stat_hit t1  INNER JOIN b_stat_country t3 ON (t3.ID = t1.COUNTRY_ID)", sql.SQL)
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
	jsonStr := `{"SELECT":["COUNTRY_ID","USER","COUNTRY"],"WHERE":"COUNTRY_ID>:countryId and SESSION_ID>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["COUNTRY_ID","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	for i := 0; i < b.N; i++ {
		NewHitSQLBuilder(filter).BuildSQL()
	}
	// Здесь выполняем функцию для тестирования
}
