package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

func main() {
	//sb := sqlbuilder.NewSelectBuilder()
	//sb.And()
	//sql := sqlbuilder.Select("id", "name").From("demo.user").
	//	Where("status = 1").Where("id=2`").
	//	String()

	fmt.Println(sqlbuilder.Select("g.id").From("guest g").Join("LEFT JOIN b_stat_country c ON (c.ID = g.LAST_COUNTRY_ID)"))

}
