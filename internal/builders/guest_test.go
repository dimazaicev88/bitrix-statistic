package builders

import (
	"bitrix-statistic/internal/filters"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGuestSQLBuilder_ToString(t *testing.T) {
	req := require.New(t)
	var tf = filters.Filter{
		Fields:    []string{"id", "last_country_name", "last_city_name"},
		Skip:      0,
		Limit:     0,
		OrderBy:   "",
		Order:     "",
		Operators: nil,
	}
	b := NewGuestBuilder(tf)
	result, err := b.ToString()
	req.Nil(err)
	req.Equal("SELECT g.id as id, c.name as last_country_name, city.last_city_name as last_city_name FROM guest g", result)
}

func BenchmarkInt(b *testing.B) {
	var tf interface{}
	tf = 10
	for i := 0; i < b.N; i++ {
		_ = tf.(int)
	}
}
