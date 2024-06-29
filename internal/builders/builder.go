package builders

import (
	"strings"
)

// TODO добавить проверку что параметры указаны правильно в where, order by, select,params.
type (
	SelectBuild struct {
		sqlData SQLDataForBuild
	}

	WhereBuilder struct {
		sqlData SQLDataForBuild
	}

	OrderByBuilder struct {
		sqlData SQLDataForBuild
	}

	SQLBuild struct {
		sqlData SQLDataForBuild
	}

	DefaultSQLBuild interface {
		SQLDataForBuild
	}

	SQL struct {
		SQL    string
		Params []interface{}
	}

	SQLDataForBuild struct {
		selectBuilder  *strings.Builder
		joinBuilder    *strings.Builder
		whereBuilder   *strings.Builder
		orderByBuilder *strings.Builder
		//filter         filters.Filter
		params *[]interface{}
		limit  *strings.Builder
		offset *strings.Builder
	}
)

//func NewSQLBuilder(filter filters.Filter) SQLDataForBuild {
//	return SQLDataForBuild{
//		selectBuilder:  &strings.Builder{},
//		joinBuilder:    &strings.Builder{},
//		whereBuilder:   &strings.Builder{},
//		orderByBuilder: &strings.Builder{},
//		limit:          &strings.Builder{},
//		offset:         &strings.Builder{},
//		filter:         filter,
//		params:         &[]interface{}{},
//	}
//}

//func NewSelectBuild(sqlData SQLDataForBuild) SelectBuild {
//	return SelectBuild{sqlData: sqlData}
//}
//
//func NewWhereBuilder(sqlData SQLDataForBuild) WhereBuilder {
//	return WhereBuilder{sqlData: sqlData}
//}
//
//func NewOrderByBuilder(sqlData SQLDataForBuild) OrderByBuilder {
//	return OrderByBuilder{sqlData: sqlData}
//}
//
//func NewSQLBuild(sqlData SQLDataForBuild) SQLBuild {
//	return SQLBuild{sqlData}
//}
//
//func (sb SelectBuild) Build(selectBuild func(sqlData SQLDataForBuild) (WhereBuilder, error)) (WhereBuilder, error) {
//	result, err := selectBuild(sb.sqlData)
//	if err != nil {
//		return WhereBuilder{}, err
//	}
//	return result, nil
//}
//
//func (wb WhereBuilder) Build(whereBuild func(sqlData SQLDataForBuild) OrderByBuilder) OrderByBuilder {
//	return whereBuild(wb.sqlData)
//}

//func (wb WhereBuilder) BuildDefault() OrderByBuilder {
//	where := wb.sqlData.filter.Where
//	if utf8.RuneCountInString(where) == 0 {
//		return NewOrderByBuilder(wb.sqlData)
//	}
//	for key, value := range wb.sqlData.filter.Params {
//		where = strings.ReplaceAll(where, key, " ? ")
//		*wb.sqlData.params = append(*wb.sqlData.params, value)
//	}
//	for key, value := range hitFields {
//		var re = regexp.MustCompile(`\b` + key + `\b`)
//		where = re.ReplaceAllString(where, value)
//	}
//	wb.sqlData.whereBuilder.WriteString(" WHERE ")
//	wb.sqlData.whereBuilder.WriteString(where)
//	return NewOrderByBuilder(wb.sqlData)
//}

//func (ob OrderByBuilder) BuildDefault() SQLBuild {
//	if len(ob.sqlData.filter.OrderBy) == 0 {
//		return NewSQLBuild(ob.sqlData)
//	}
//	ob.sqlData.orderByBuilder.WriteString(" ORDER BY ")
//	var orderByFields []string
//	for _, by := range ob.sqlData.filter.OrderBy {
//		orderByFields = append(orderByFields, hitFields[by])
//	}
//	ob.sqlData.orderByBuilder.WriteString(strings.Join(orderByFields, ","))
//	ob.sqlData.orderByBuilder.WriteString(" ")
//	ob.sqlData.orderByBuilder.WriteString(ob.sqlData.filter.TypeSort)
//	return NewSQLBuild(ob.sqlData)
//}
//
//func (sb SQLBuild) DefaultBuild(selectBuilder func() (WhereBuilder, error)) (SQL, error) {
//	var resultSQL strings.Builder
//	sqlBuilder, err := selectBuilder()
//	if err != nil {
//		return SQL{}, err
//	}
//	sqlBuilder.BuildDefault().
//		BuildDefault()
//	resultSQL.WriteString(sb.sqlData.selectBuilder.String())
//	resultSQL.WriteString(sb.sqlData.whereBuilder.String())
//	resultSQL.WriteString(sb.sqlData.orderByBuilder.String())
//	return SQL{
//		SQL:    resultSQL.String(),
//		Params: *sb.sqlData.params,
//	}, nil
//}
