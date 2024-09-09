package builders

var pathFields = map[string]string{
	"PATH_ID":            " t1.PATH_ID",             // ID отрезка пути
	"DATE1":              " t1.DATE1 ",              // начальное значение для интервала даты
	"DATE2":              " t1.DATE2 ",              // конечное значение для интервала даты
	"FIRST_PAGE":         " t1.FIRST_PAGE ",         // первая страница пути
	"FIRST_PAGE_SITE_ID": " t1.FIRST_PAGE_SITE_ID ", // ID сайта первой страницы пути
	"FIRST_PAGE_404":     " t1.FIRST_PAGE_404 ",     // была ли 404 ошибка на первой странице пути, возможные значения:Y - была,	N - не была.
	"LAST_PAGE":          " t1.LAST_PAGE ",          //последняя страница пути
	"LAST_PAGE_SITE_ID":  " t1.LAST_PAGE_SITE_ID ",  // ID, сайта последней страницы пути была ли 404 ошибка на последней странице пути, возможные значения: Y - была, N - не была.
	"LAST_PAGE_404":      " t1.LAST_PAGE_404 ",      // была ли 404 ошибка на последней странице пути, возможные значения: Y - была; N - не была.
	"PAGE":               " t1.PAGE ",               //произвольная страница пути
	"PAGE_SITE_ID":       " t1.PAGE_SITE_ID ",       //ID сайта произвольной страницы пути
	"PAGE_404":           " t1.PAGE_404 ",           //была ли 404 ошибка на произвольной странице пути, возможные значения:    Y - была    N - не была.
	"ADV_ID":             " t2.ADV_ID ",             //ID рекламной кампании, по посетителям которой надо получить данные
	"ADV_DATA_TYPE":      " t2.ADV_DATA_TYPE ",      //флаг типа данных для рекламной кампании, возможные значения:
	//P - только по прямым заходам по рекламной кампании
	//B - только по возвратам по рекламной кампании
	//S - сумма по прямым заходам и возвратам.
	"STEPS1": " t1.CountryId ", //начальное значение интервала для поля "количество страниц в пути"
	"STEPS2": " t1.CityUuid ",  //конечное значение интервала для поля "количество страниц в пути".
}

//func (hs PathSqlBuilder) buildSelectAndJoin() (WhereBuilder, error) {
//	return NewSelectBuild(hs.sqlData).build(func(sqlData SQLDataForBuild) (WhereBuilder, error) {
//		var selectBuffer []string
//		sqlData.selectBuilder.WriteString("SELECT ")
//		if len(sqlData.filter.Select) == 0 {
//			sqlData.selectBuilder.WriteString(" * ")
//		} else {
//			for _, selectField := range sqlData.filter.Select {
//				if value, ok := hitSelectFields[selectField]; ok {
//					selectBuffer = append(selectBuffer, value)
//				} else {
//					return WhereBuilder{}, errors.New("unknown field " + selectField)
//				}
//				if selectField == "USER" {
//					sqlData.selectBuilder.WriteString("")
//					sqlData.joinBuilder.WriteString(" LEFT JOIN b_user t2 ON (t2.ID = t1.USER_ID)")
//				}
//				if selectField == "Country" {
//					sqlData.joinBuilder.WriteString(" INNER JOIN b_stat_country t3 ON (t3.ID = t1.CountryId)")
//				}
//			}
//		}
//		hs.sqlData.selectBuilder.WriteString(strings.Join(selectBuffer, ","))
//		hs.sqlData.selectBuilder.WriteString(" FROM b_stat_hit t1 ")
//		hs.sqlData.selectBuilder.WriteString(sqlData.joinBuilder.String())
//		return NewWhereBuilder(sqlData), nil
//	})
//}
