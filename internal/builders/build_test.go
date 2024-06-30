package builders

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"testing"
)

var testFilter = filters.Filter{
	Fields:         nil,
	Skip:           0,
	Limit:          0,
	OrderBy:        "",
	Order:          "",
	FilterOperator: nil,
}

func TestBuildWhereSQLCase1(t *testing.T) {
	req := require.New(t)
	testFilter.FilterOperator = []filters.FilterOperator{
		{
			Operator: "=",
			Value:    1,
			Field:    "id",
		},
	}
	sql, args, err := BuildWhereSQL(testFilter, func(field string) bool {
		return true
	})
	req.NoError(err)
	req.Equal(sql, "where id = ?")
	req.Equal(args, []interface{}{1})

}

func TestBuildWhereSQLCase2(t *testing.T) {
	req := require.New(t)
	testFilter.FilterOperator = []filters.FilterOperator{
		{
			Operator: "=",
			Value:    1,
			Field:    "id",
		},
		{
			Operator: "or",
			Value:    "",
			Field:    "",
		},
		{
			Operator: "=",
			Value:    10,
			Field:    "id",
		},
	}

	sql, args, err := BuildWhereSQL(testFilter, func(field string) bool {
		return true
	})
	req.NoError(err)
	req.Equal(sql, "where id = ? or id = ?")
	req.Equal(args, []interface{}{1, 10})
}

func TestBuildWhereSQLCase3(t *testing.T) {
	req := require.New(t)
	testFilter.FilterOperator = []filters.FilterOperator{
		{
			Operator: ">",
			Value:    1,
			Field:    "id",
		},
		{
			Operator: "and",
			Value:    "",
			Field:    "",
		},
		{
			Operator: "<",
			Value:    10,
			Field:    "id",
		},
	}

	sql, args, err := BuildWhereSQL(testFilter, func(field string) bool {
		return true
	})
	req.NoError(err)
	req.Equal(sql, "where id > ? and id < ?")
	req.Equal(args, []interface{}{1, 10})
}

func TestBuildWhereSQLCase4(t *testing.T) {
	req := require.New(t)
	testFilter.FilterOperator = []filters.FilterOperator{
		{
			Operator: ">",
			Value:    1,
			Field:    "id",
		},
		{
			Operator: "and",
			Value:    "",
			Field:    "",
		},
		{
			Operator: "<",
			Value:    10,
			Field:    "id",
		},
		{
			Operator: "or",
			Value:    "",
			Field:    "",
		},
		{
			Operator: ">",
			Value:    100,
			Field:    "id",
		},
	}

	sql, args, err := BuildWhereSQL(testFilter, func(field string) bool {
		return true
	})
	req.NoError(err)
	req.Equal(sql, "where id > ? and id < ? or id > ?")
	req.Equal(args, []interface{}{1, 10, 100})
}
