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
	for _, value := range filter.FilterOperator {
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
