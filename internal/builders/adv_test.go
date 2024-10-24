package builders

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdvBuilder_FilterAdv(t *testing.T) {
	req := require.New(t)
	advSQLBuilder := NewAdvSQLBuilder(filters.Filter{
		Fields:    advSelectFields,
		Skip:      10,
		Limit:     100,
		OrderBy:   "desc",
		Order:     []string{"uuid"},
		Operators: nil,
	})

	t.Run("Указано 'Куда пришли'", func(t *testing.T) {
		sql, args, err := advSQLBuilder.Build()
		req.NoError(err)
		req.Equal(args, []any{10, 0})
		req.Equal(sql, "")
	})
}
