package builders

var cityFields = map[string]string{
	"ID":         "t1.ID",
	"SHORT_NAME": "t1.SHORT_NAME",
	"NAME":       "t1.NAME",
	"SESSIONS":   "t1.SESSIONS",
	"NEW_GUESTS": "t1.NEW_GUESTS",
	"HITS":       "t1.HITS",
	"C_EVENTS":   "t1.C_EVENTS",
}

//type CountrySQLBuilder struct {
//	sqlData SQLDataForBuild
//}

//func NewCountrySQLBuilder(filter filters.Filter) CountrySQLBuilder {
//	return CountrySQLBuilder{NewSQLBuilder(filter)}
//}

//func (cb CountrySQLBuilder) buildSelectAndGroupBy() (WhereBuilder, error) {
//return NewSelectBuild(cb.sqlData).build(func(sqlData SQLDataForBuild) (WhereBuilder, error) {
//	var selectBuffer []string
//	cb.sqlData.selectBuilder.WriteString("SELECT ")
//	if len(cb.sqlData.filter.Select) == 0 {
//		cb.sqlData.selectBuilder.WriteString("* ")
//	} else {
//		for _, selectField := range cb.sqlData.filter.Select {
//			if value, ok := cityFields[selectField]; ok {
//				selectBuffer = append(selectBuffer, value)
//			}
//		}
//	}
//	if len(selectBuffer) == 0 {
//		return WhereBuilder{}, errors.New("unknown fields is select")
//	}
//	cb.sqlData.selectBuilder.WriteString(strings.Join(selectBuffer, ","))
//	cb.sqlData.selectBuilder.WriteString(" FROM b_stat_country t1 ")
//	cb.sqlData.selectBuilder.WriteString(cb.sqlData.joinBuilder.String())
//	return NewWhereBuilder(sqlData), nil
//})
//}

//func (cb CountrySQLBuilder) orderByBuild() SQLBuild {
//	//return NewOrderByBuilder(cb.sqlData).BuildDefault()
//}

// TODO добавить проверку на наличие неизвестных полей.
//func (cb CountrySQLBuilder) whereBuild() OrderByBuilder {
//	return NewWhereBuilder(cb.sqlData).BuildDefault()
//}

//func (cb CountrySQLBuilder) BuildSQL() (SQL, error) {
//	return NewSQLBuild(cb.sqlData).DefaultBuild(cb.buildSelectAndGroupBy)
//}
