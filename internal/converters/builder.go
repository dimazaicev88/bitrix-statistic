package converters

import "strings"

type FilterToSqlConverter struct {
	resultSql []string
	args      []any
}

func NewSqlSQLConverter() *FilterToSqlConverter {
	return &FilterToSqlConverter{
		resultSql: make([]string, 0, 10),
		args:      []any{},
	}
}

func (sb *FilterToSqlConverter) AddSql(sql string) *FilterToSqlConverter {
	sb.resultSql = append(sb.resultSql, sql)
	return sb
}

func (sb *FilterToSqlConverter) AddArgs(args ...any) *FilterToSqlConverter {
	sb.args = append(sb.args, args...)
	return sb
}

func (sb *FilterToSqlConverter) Convert() (string, []any) {
	return strings.Join(sb.resultSql, " "), sb.args
}
