package main

import (
	"bitrix-statistic/internal/builders"
	"bitrix-statistic/internal/filters"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

func main() {
	var filter filters.Filter
	jsonStr := `{"SELECT":["COUNTRY_ID","USER","COUNTRY"],"WHERE":"COUNTRY_ID>:countryId and SESSION_ID>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["COUNTRY_ID","ID"],"TYPE_SORT":"ASC"}`
	jsoniter.Unmarshal([]byte(jsonStr), &filter)
	start := time.Now()
	fmt.Println(builders.NewHitSQLBuilder(filter).BuildSQL())

	stop := time.Now().Sub(start)
	fmt.Println(stop)
}
