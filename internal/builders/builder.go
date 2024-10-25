package builders

import "strings"

type SqlBuilder struct {
	resultSql []string
	args      []any
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{
		resultSql: make([]string, 0, 10),
		args:      []any{},
	}
}

func (sb *SqlBuilder) AddSql(sql string) *SqlBuilder {
	sb.resultSql = append(sb.resultSql, sql)
	return sb
}

func (sb *SqlBuilder) AddArgs(args ...any) *SqlBuilder {
	sb.args = append(sb.args, args...)
	return sb
}

func (sb *SqlBuilder) Build() (string, []any) {
	return strings.Join(sb.resultSql, " "), sb.args
}
