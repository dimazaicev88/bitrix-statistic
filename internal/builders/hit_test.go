package builders

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHitSQLBuilder(t *testing.T) {
	req := require.New(t)
	testFilter.Operators = []filters.Operators{
		{
			Operator: "=",
			Value:    1,
			Field:    "id",
		},
	}
	sql, args, err := BuildWhereSQL(testFilter)
	req.NoError(err)
	req.Equal(sql, "where id = ?")
	req.Equal(args, []interface{}{1})
}
