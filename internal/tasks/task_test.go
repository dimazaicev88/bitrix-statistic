package tasks

import (
	"bitrix-statistic/internal/cache"
	"testing"
)

func BenchmarkAddTask(b *testing.B) {

	for i := 0; i < b.N; i++ {
		cache.AdvDays()
	}
}
