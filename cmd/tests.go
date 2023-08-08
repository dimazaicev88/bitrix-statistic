package main

//func main() {
//	var filter filters.Filter
//	jsonStr := `{"SELECT":["USER","Country"],"WHERE":"CountryId>:countryId and SessionId>:sessionId","PARAMS":{":countryId":1,":sessionId":12},"ORDER_BY":["CountryId","ID"],"TYPE_SORT":"ASC"}`
//	jsoniter.Unmarshal([]byte(jsonStr), &filter)
//	start := time.Now()
//	fmt.Println(builders.NewHitSQLBuilder(filter).BuildSQL())
//
//	stop := time.Now().Sub(start)
//	fmt.Println(stop)
//
//	//set := utils.NewSet[string]()
//	//set.Add("qwe")
//}
