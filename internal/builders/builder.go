package builders

import (
	"bitrix-statistic/internal/filters"
	"errors"
	"slices"
	"strings"
)

var allOperator = []string{"=", ">=", "=<", ">", "<", "!=", "<>", "like", "not like", "and", "or"}

func BuildWhereSQL(filter filters.Filter, validWhereField func(field string) bool) (string, []interface{}, error) {
	var strBuilder strings.Builder
	var args []interface{}
	strBuilder.WriteString("where ")
	for _, value := range filter.Operators {
		if slices.Contains(allOperator, value.Operator) == false {
			return "", nil, errors.New("unknown operator: " + value.Operator)
		}

		if !validWhereField(value.Field) {
			return "", nil, errors.New("unknown field: " + value.Field)
		}

		if len(strBuilder.String()) == 0 && (value.Operator == "and" || value.Operator == "or" || value.Operator == "like" || value.Operator == "not like") {
			return "", nil, errors.New("invalid position operator. where " + value.Operator)
		}
		strBuilder.WriteString(value.Field)
		strBuilder.WriteString(" ")
		strBuilder.WriteString(value.Operator)
		strBuilder.WriteString(" ")
		if value.Operator != "and" && value.Operator != "or" {
			strBuilder.WriteString("?")
			args = append(args, value.Value)
		}
	}

	return strBuilder.String(), args, nil
}

func BuildLimit(filter filters.Filter) (string, []int) {
	var strBuilder strings.Builder
	var args []int
	strBuilder.WriteString("limit ")
	if filter.Skip != 0 {
		args = append(args, filter.Skip)
		strBuilder.WriteString("?")
		strBuilder.WriteString(", ")
	}

	if filter.Limit != 0 && filter.Limit < 1000 {
		args = append(args, filter.Limit)
		strBuilder.WriteString("?")
	} else {
		args = append(args, 1000)
		strBuilder.WriteString("?")
	}
	return strBuilder.String(), args
}

func BuildOrderBy() {

}
