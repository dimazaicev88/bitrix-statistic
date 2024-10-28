package converters

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHitSQLBuilder(t *testing.T) {
	req := require.New(t)
	filter := filters.Filter{
		Fields:  []string{"uuid", "sessionUuid", "advUuid", "dateHit", "phpSessionId"},
		Skip:    0,
		Limit:   0,
		OrderBy: "",
		Order:   nil,
		Operators: []filters.Operators{
			{
				Operator: "=",
				Value:    1,
				Field:    "uuid",
			},
			{
				Operator: "=",
				Value:    true,
				Field:    "isRegistered",
			},
			{
				Operator: "or",
				Value:    nil,
				Field:    "",
			},
			{
				Operator: "=",
				Value:    true,
				Field:    "isNewUser",
			},
		},
	}
	hitBuilder := NewHitSQLBuilder(filter)
	sql, args, err := hitBuilder.Build()
	if err != nil {
		return
	}
	req.NoError(err)
	req.Equal(sql, "SELECT uuid, sessionUuid, advUuid, dateHit, phpSessionId FROM hit WHERE uuid=? AND userId>0  OR isNewUser=? LIMIT ? ?")
	req.Equal(args, []any{1, true, 0, 1000})
}
