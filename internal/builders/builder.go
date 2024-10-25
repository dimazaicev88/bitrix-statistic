package builders

import "strings"

type SqlBuilder struct {
	resultSql strings.Builder
	args      []any
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{
		resultSql: strings.Builder{},
		args:      []any{},
	}
}

func (sb *SqlBuilder) Add(str string, args ...any) *SqlBuilder {
	sb.resultSql.WriteString(str)
	sb.args = append(sb.args, args...)
	return sb
}

func (sb *SqlBuilder) Build() (string, []any) {
	return sb.resultSql.String(), sb.args
}
