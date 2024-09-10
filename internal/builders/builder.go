package builders

import "strings"

type SqlBuilder struct {
	resultSql strings.Builder
	args      []interface{}
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{
		resultSql: strings.Builder{},
		args:      []interface{}{},
	}
}

func (sb *SqlBuilder) Add(str string, args ...interface{}) *SqlBuilder {
	sb.resultSql.WriteString(str)
	sb.args = append(sb.args, args...)
	return sb
}

func (sb *SqlBuilder) Build() (string, []interface{}) {
	return sb.resultSql.String(), sb.args
}
